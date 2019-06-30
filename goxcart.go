package goxcart

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// Version is semver.
var Version = "0.0.4"

// ContainerIDPattern matches Docker container IDs as reported by `docker run -d`...
var ContainerIDPattern = regexp.MustCompile(`^(?P<container>[0-9a-f]+)\s+$`)

// SuccessExitStatusPattern matches successful Docker wait output.
var SuccessExitStatusPattern = regexp.MustCompile("^0$")

// // atomicKilled marks whether the goxcart process is being killed.
// var atomicKilled AtomicKilled

// daemonLock prevents Docker daemons from terminating early when goxcart termintates.
var daemonLock sync.Mutex

// atomicContainerIDs marks active Docker container IDs.
var atomicContainerIDs = NewAtomicContainerIDs()

// PortPlatformGroup builds application ports for a specific platform group.
func (o PortConfig) portJob(job PortJob) error {
	daemonLock.Lock()

	platformGroup := job.PlatformGroup
	cwd := job.CurrentWorkingDirectory
	env := job.Environment
	sourcePath := job.SourcePath

	if o.Verbose {
		log.Printf("PlatformGroup: %v\n", platformGroup)
	}

	var dockerImageBuffer bytes.Buffer
	dockerImageBuffer.WriteString(o.Image)

	if platformGroup.ImageTag != "" {
		dockerImageBuffer.WriteString(":")
		dockerImageBuffer.WriteString(platformGroup.ImageTag)
	}

	var goxOutputStructureBuffer bytes.Buffer
	goxOutputStructureBuffer.WriteString(o.OutputDirectory)
	goxOutputStructureBuffer.WriteString("/")
	goxOutputStructureBuffer.WriteString(o.Banner)
	goxOutputStructureBuffer.WriteString("/{{.OS}}")

	if platformGroup.OSVariant != "" {
		goxOutputStructureBuffer.WriteString("-")
		goxOutputStructureBuffer.WriteString(platformGroup.OSVariant)
	}

	goxOutputStructureBuffer.WriteString("/{{.Arch}}/{{.Dir}}")

	var dockerCommandBuffer bytes.Buffer
	dockerCommandBuffer.WriteString("cd ")
	dockerCommandBuffer.WriteString(sourcePath)
	dockerCommandBuffer.WriteString(" && gox -output='")
	dockerCommandBuffer.WriteString(goxOutputStructureBuffer.String())
	dockerCommandBuffer.WriteString("'")

	if len(platformGroup.OSs) > 0 {
		dockerCommandBuffer.WriteString(" -os='")
		dockerCommandBuffer.WriteString(strings.Join(platformGroup.OSs, " "))
		dockerCommandBuffer.WriteString("'")
	}

	if len(platformGroup.Archs) > 0 {
		dockerCommandBuffer.WriteString(" -arch='")
		dockerCommandBuffer.WriteString(strings.Join(platformGroup.Archs, " "))
		dockerCommandBuffer.WriteString("'")
	}

	if platformGroup.LinkerFlags != "" {
		dockerCommandBuffer.WriteString(" -ldflags='")
		dockerCommandBuffer.WriteString(platformGroup.LinkerFlags)
		dockerCommandBuffer.WriteString("'")
	}

	dockerCommandBuffer.WriteString(" ")
	dockerCommandBuffer.WriteString(o.CommandPaths)

	var daemonOut bytes.Buffer

	var daemonParts []string
	daemonParts = append(daemonParts, "run")

	if o.RemoveContainer {
		daemonParts = append(daemonParts, "--rm")
	}

	daemonParts = append(daemonParts, "-d")
	daemonParts = append(daemonParts, "-v")
	daemonParts = append(daemonParts, fmt.Sprintf("%s:%s", cwd, sourcePath))
	daemonParts = append(daemonParts, dockerImageBuffer.String())
	daemonParts = append(daemonParts, "sh")
	daemonParts = append(daemonParts, "-c")
	daemonParts = append(daemonParts, dockerCommandBuffer.String())

	cmdDaemon := exec.Command("docker", daemonParts...)
	cmdDaemon.Env = env
	cmdDaemon.Stdout = bufio.NewWriter(&daemonOut)
	cmdDaemon.Stderr = os.Stderr

	if o.Verbose {
		log.Printf("Docker daemon: %s", strings.Join(cmdDaemon.Args, " "))
	}

	err := cmdDaemon.Run()

	daemonLock.Unlock()

	if err != nil {
		return err
	}

	daemonOutString := daemonOut.String()

	m := ContainerIDPattern.FindStringSubmatch(daemonOutString)

	if len(m) < 2 {
		return fmt.Errorf("could not extract container ID from %s", daemonOutString)
	}

	id := m[1]

	defer atomicContainerIDs.RemoveID(id)
	atomicContainerIDs.AddID(id)

	var waitOut bytes.Buffer

	cmdBlock := exec.Command("docker", "wait", id)
	cmdBlock.Stdout = bufio.NewWriter(&waitOut)
	cmdBlock.Stderr = os.Stderr

	if o.Verbose {
		log.Printf("Waiting on container %s\n", id)
	}

	if err := cmdBlock.Run(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(&waitOut)

	for scanner.Scan() {
		status := scanner.Text()

		if !SuccessExitStatusPattern.MatchString(status) {
			return fmt.Errorf("docker wait %s returned exit status %s", id, status)
		}
	}

	return nil
}

// worker listens for jobs to process.
func (o *PortConfig) worker(jobCh <-chan PortJob, wg *sync.WaitGroup, atomicErrors *AtomicErrors) {
	for job := range jobCh {
		if err := o.portJob(job); err != nil {
			atomicErrors.AddError(err)
		}

		wg.Done()
	}
}

// Port builds application ports.
func (o *PortConfig) Port() error {
	if err := o.Validate(); err != nil {
		return err
	}

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	if o.Repository == "" {
		o.Repository = path.Base(cwd)
	}

	sourcePath := fmt.Sprintf("/go/src/%s", o.Repository)
	env := os.Environ()

	var wg sync.WaitGroup
	wg.Add(len(PlatformGroups))

	var atomicErrors AtomicErrors

	jobCh := make(chan PortJob)
	defer close(jobCh)

	for i := 0; i < runtime.NumCPU(); i++ {
		go o.worker(jobCh, &wg, &atomicErrors)
	}

	for _, platformGroup := range PlatformGroups {
		portJob := PortJob{
			PlatformGroup:           platformGroup,
			CurrentWorkingDirectory: cwd,
			Environment:             env,
			SourcePath:              sourcePath,
		}

		jobCh <- portJob
	}

	wg.Wait()

	errs := atomicErrors.GetErrors()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

// KillContainers terminates goxcart's active Docker containers.
func KillContainers() error {
	defer daemonLock.Unlock()
	daemonLock.Lock()

	ids := atomicContainerIDs.GetIDs()

	for _, id := range ids {
		cmd := exec.Command("docker", "kill", id)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
