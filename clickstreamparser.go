package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Command string

const (
	R_AD        Command = "41"
	R_BtnCnfg   Command = "42"
	R_ChanVrb   Command = "43"
	R_STATE     Command = "53"
	R_HIGHLIGHT Command = "48"
	R_INFO      Command = "49"
	R_VIDEO     Command = "56"
	R_KEY       Command = "4B"
	R_UNIT      Command = "55"
)

func CheckCommand(clickString string) Command {
	return Command(clickString[0:2])
}

const (
	txtOutput  = "txt"
	xmlOutput  = "xml"
	jsonOutput = "json"
	rawExt     = "raw"
)

func init() {
	flagFileName := flag.String("fi", "", "Single filename to process")
	flagDirName := flag.String("d", "", "Working folder for input files, default extension *.raw")
	flagExtension := flag.String("e", rawExt, "Input files extension")
	flagDiagnostics := flag.Bool("t", false, "Turns diagnostic messages On")
	flagOutputFormat := flag.String("o", txtOutput, "Output formats: txt, json, xml")
	flagOutputFile := flag.String("fo", "output", "Output file name")
	flagConcurrency := flag.Int("c", 100, "The number of files to process concurrently")
	flagVerbose := flag.Bool("v", true, "Verbose: outputs to the screen")

	flag.Parse()
	if flag.Parsed() {
		inFileName = *flagFileName
		dirName = *flagDirName
		inExtension = *flagExtension
		diagnostics = *flagDiagnostics
		outputFormat = *flagOutputFormat
		outputFileName = *flagOutputFile
		concurrency = *flagConcurrency
		verbose = *flagVerbose
	} else {
		flag.Usage()
		os.Exit(-1)
	}

}

var (
	inFileName     string
	dirName        string
	inExtension    string
	diagnostics    bool
	outputFormat   string
	outputFileName string
	concurrency    int
	verbose        bool
	singleFileMode bool
)

func main() {
	startTime := time.Now()

	// This is our semaphore/pool
	sem := make(chan bool, concurrency)
	totalEventsChan := make(chan int, concurrency)

	files := getFilesToProcess()

	totalEvents := 0
	go func() {
		for {
			nn, more := <-totalEventsChan
			if more {
				if diagnostics {
					fmt.Println("Reported: ", nn)
				}
				totalEvents += nn
			} else {
				if diagnostics {
					fmt.Println("Got all reports, breaking")
				}
				return
			}
		}
	}()

	for _, gfile := range files {
		// if we still have available goroutine in the pool (out of concurrency )
		sem <- true

		// fire one file to be processed in a goroutine
		go func(fileName string) {
			// Signal end of processing at the end
			defer func() { <-sem }()
			eventsCollection := []interface{}{}

			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Error opening file: ", err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				tokens := strings.Split(line, " ")
				if len(tokens) != 2 {
					fmt.Println("Wrong file format for: ", fileName)
					return
				}
				deviceId, clickString := tokens[0], tokens[1]
				switch CheckCommand(clickString) {
				case R_AD:
					adEvent := NewAdEvent(deviceId, clickString)
					if verbose {
						fmt.Println(adEvent)
					}
					eventsCollection = append(eventsCollection, adEvent)
					if diagnostics {
						fmt.Println("Diagnostics: ", adEvent.BaseEvent.Diagnostic())

						fmt.Println(adEvent.Command,
							adEvent.Timestamp,
							adEvent.AdType,
							adEvent.AdId,
							adEvent.Serial,
							adEvent.Checksum,
							adEvent.Linefeed)
					}
				case R_BtnCnfg:
					btcnfgEvent := NewButtonConfigEvent(deviceId, clickString)
					if verbose {
						fmt.Println(btcnfgEvent)
					}
					eventsCollection = append(eventsCollection, btcnfgEvent)
					if diagnostics {
						fmt.Println("Diagnostics: ", btcnfgEvent.BaseEvent.Diagnostic())
						fmt.Println(btcnfgEvent.Command,
							btcnfgEvent.Timestamp,
							btcnfgEvent.ButtonId,
							btcnfgEvent.ButtonType,
							btcnfgEvent.ButtonText,
							btcnfgEvent.ButtonVarData,
							btcnfgEvent.Serial,
							btcnfgEvent.Checksum,
							btcnfgEvent.Linefeed)
					}
				case R_ChanVrb:
					channelchange := NewChannelChangeVerboseEvent(deviceId, clickString)
					if verbose {
						fmt.Println(channelchange)
					}
					eventsCollection = append(eventsCollection, channelchange)
					if diagnostics {
						fmt.Println("Diagnostics: ", channelchange.BaseEvent.Diagnostic())
						fmt.Println(channelchange.Command,
							channelchange.Timestamp,
							channelchange.Channel,
							channelchange.SourseId,
							channelchange.ProgramId,
							channelchange.Auth,
							channelchange.TunerInfo,
							channelchange.PreviousState,
							channelchange.LastKey,
							channelchange.Serial,
							channelchange.Checksum,
							channelchange.Linefeed)
					}
				case R_STATE:
					statechange := NewStateEvent(deviceId, clickString)
					if verbose {
						fmt.Println(statechange)
					}
					eventsCollection = append(eventsCollection, statechange)
					if diagnostics {
						fmt.Println("Diagnostics: ", statechange.BaseEvent.Diagnostic())
						fmt.Println(statechange.Command,
							statechange.State,
							statechange.PreviousState,
							statechange.LastKey)
					}
				case R_INFO:
					info := NewInfoScreenEvent(deviceId, clickString)
					if verbose {
						fmt.Println(info)
					}
					eventsCollection = append(eventsCollection, info)
					if diagnostics {

						fmt.Println("Diagnostics: ", info.BaseEvent.Diagnostic())
						fmt.Println(info.Command,
							info.Type,
							info.Id)
					}
				case R_KEY:
					key := NewKeyPressEvent(deviceId, clickString)
					if verbose {
						fmt.Println(key)
					}
					eventsCollection = append(eventsCollection, key)
					if diagnostics {

						fmt.Println("Diagnostics: ", key.BaseEvent.Diagnostic())
						fmt.Println(key.Command,
							key.KeyCode)
					}
				case R_HIGHLIGHT:
					hilit := NewHighlightEvent(deviceId, clickString)
					if verbose {
						fmt.Println(hilit)
					}
					eventsCollection = append(eventsCollection, hilit)
					if diagnostics {

						fmt.Println("Diagnostics: ", hilit.BaseEvent.Diagnostic())
						fmt.Println(hilit.Command,
							hilit.Type,
							hilit.IdFieldsStr)
					}
				case R_VIDEO:
					video := NewVideoPlaybackEvent(deviceId, clickString)
					if verbose {
						fmt.Println(video)
					}
					eventsCollection = append(eventsCollection, video)
					if diagnostics {

						fmt.Println("Diagnostics: ", video.BaseEvent.Diagnostic())
						fmt.Println(video.Id,
							video.VodPlaybackMode,
							video.Source,
							video.PlayBackPosition)
					}
				case R_UNIT:
					unit := NewUnitIdentificationEvent(deviceId, clickString)
					if verbose {
						fmt.Println(unit)
					}
					eventsCollection = append(eventsCollection, unit)
					if diagnostics {

						fmt.Println("Diagnostics: ", unit.BaseEvent.Diagnostic())
						fmt.Println(unit.PeriodicReports,
							unit.PollingReports,
							unit.HighWaterMarkReports,
							unit.BlackoutOverflowReports,
							unit.ExceededMaxReportsPerHour,
							unit.UsedBufferSize,
							unit.GuideState,
							unit.TunerInfo,
							unit.SourceIdTuner0,
							unit.SourceIdTuner1)
					}
				}
			}

			if err = scanner.Err(); err != nil {
				fmt.Printf("Error while processing file: %s: %v\n", fileName, err)
			}
			// Reporting number of processed events
			totalEventsChan <- len(eventsCollection)

			fileNameToSave := formatFileNameToSave(fileName)
			switch outputFormat {
			case jsonOutput:
				processJson(fileNameToSave, eventsCollection)
			case xmlOutput:
				processXml(fileNameToSave, eventsCollection)
			case txtOutput:
				processText(fileNameToSave, eventsCollection)
			}
		}(gfile)
	}

	// waiting for all goroutines to end
	if diagnostics {
		fmt.Println("Waiting for all goroutines to complete the work")
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	// Done all gouroutines, close the total channel
	if diagnostics {
		fmt.Println("Closing total events channel")
	}

	close(totalEventsChan)

	fmt.Printf("Processed %d files, %d events in %v\n", len(files), totalEvents, time.Since(startTime))

}

// Format output file name
func formatFileNameToSave(currentFileName string) string {
	if singleFileMode {
		return validateOutFileName(outputFileName)
	}
	return validateOutFileName(currentFileName[:len(currentFileName)-len("."+inExtension)])
}

// Get the list of files to process in the target folder
func getFilesToProcess() []string {
	fileList := []string{}
	singleFileMode = false

	if dirName == "" {
		if inFileName != "" {
			// no Dir name provided, but file name provided =>
			// Single file mode
			singleFileMode = true
			fileList = append(fileList, inFileName)
			return fileList
		} else {
			// no Dir name, no file name
			fmt.Println("Input file name or working directory is not provided")
			flag.Usage()
			os.Exit(-1)
		}
	}

	// We have working directory - takes over single file name, if both provided
	err := filepath.Walk(dirName, func(path string, f os.FileInfo, _ error) error {
		if isRawFile(path) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error getting files list: ", err)
		os.Exit(-1)
	}

	return fileList
}

func processText(filename string, eventsCollection []interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(file)
	for _, event := range eventsCollection {
		fmt.Fprintln(w, event)
	}
	w.Flush()
	file.Close()
}

func processJson(filename string, eventsCollection []interface{}) {
	jsonString, err := generateJson(eventsCollection)
	if diagnostics {
		fmt.Println(string(jsonString))
	}
	if err == nil {
		err = saveJsonToFile(filename, jsonString)
		if err != nil {
			fmt.Println("Error writing Json file:", err)
		}
	} else {
		fmt.Println(err)
	}

}

func processXml(filename string, eventsCollection []interface{}) {
	xmlString, err := generateXml(eventsCollection)
	if diagnostics {
		fmt.Println(string(xmlString))
	}
	if err == nil {
		err = saveXmlToFile(filename, xmlString)
		if err != nil {
			fmt.Println("Error writing XML file:", err)
		}
	} else {
		fmt.Println(err)
	}

}
