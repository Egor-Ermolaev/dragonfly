package cmd

// Source represents a source of a command execution. Commands may limit the sources that can run them by
// implementing the Allower interface.
// Source implements Target. A Source must always be able to target itself.
type Source interface {
	Target
	// SendCommandOutput sends a command output to the source. The way the output is applied, depends on what
	// kind of source it is.
	// SendCommandOutput is called by a Command automatically after being run.
	SendCommandOutput(o *Output)
}
