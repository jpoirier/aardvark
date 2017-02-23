// Copyright (c) 2017 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// Package aardvark wraps aardvark.so/.dll

package aardvark

import (
	"errors"
	// "fmt"
	"unsafe"
)

/*

#cgo !windows LDFLAGS: -L/usr/local/lib64 -laardvark
#cgo windows LDFLAGS: -laardvark -LC:/WINDOWS/system32

#include <stdlib.h>
#include <stdio.h>
#include <aardvark.h>
*/
import "C"

// Version
const Version = "0.1.0"

// Device
type Device struct {
	aa C.Aardvark
	fp C.int
}

// AardvarkStatus
type AardvarkStatus = C.AardvarkStatus

var (
	// General codes (0 to -99)
	errNil                  AardvarkStatus = C.AA_OK
	errUnableToLoadLibrary  AardvarkStatus = C.AA_UNABLE_TO_LOAD_LIBRARY
	errUnableToLoadDriver   AardvarkStatus = C.AA_UNABLE_TO_LOAD_DRIVER
	errUnableToLoadFunction AardvarkStatus = C.AA_UNABLE_TO_LOAD_FUNCTION
	errIncompatibleLibrary  AardvarkStatus = C.AA_INCOMPATIBLE_LIBRARY
	errIncompatibleDevice   AardvarkStatus = C.AA_INCOMPATIBLE_DEVICE
	errCommunicationError   AardvarkStatus = C.AA_COMMUNICATION_ERROR
	errUnableToOpen         AardvarkStatus = C.AA_UNABLE_TO_OPEN
	errUnableToClose        AardvarkStatus = C.AA_UNABLE_TO_CLOSE
	errInvalidHandle        AardvarkStatus = C.AA_INVALID_HANDLE
	errConfigError          AardvarkStatus = C.AA_CONFIG_ERROR

	// I2C codes (-100 to -199)
	errI2cNotAvailable       AardvarkStatus = C.AA_I2C_NOT_AVAILABLE
	errI2cNotEnabled         AardvarkStatus = C.AA_I2C_NOT_ENABLED
	errI2cReadError          AardvarkStatus = C.AA_I2C_READ_ERROR
	errI2cWriteError         AardvarkStatus = C.AA_I2C_WRITE_ERROR
	errI2cSlaveBadConfig     AardvarkStatus = C.AA_I2C_SLAVE_BAD_CONFIG
	errI2cSlaveReadError     AardvarkStatus = C.AA_I2C_SLAVE_READ_ERROR
	errI2cSlaveTimeout       AardvarkStatus = C.AA_I2C_SLAVE_TIMEOUT
	errI2cDroppedExcessBytes AardvarkStatus = C.AA_I2C_DROPPED_EXCESS_BYTES
	errI2cBusAlreadyFree     AardvarkStatus = C.AA_I2C_BUS_ALREADY_FREE

	// SPI codes (-200 to -299)
	errSpiNotAvailable       AardvarkStatus = C.AA_SPI_NOT_AVAILABLE
	errSpiNotEnabled         AardvarkStatus = C.AA_SPI_NOT_ENABLED
	errSpiWriteError         AardvarkStatus = C.AA_SPI_WRITE_ERROR
	errSpiSlaveReadError     AardvarkStatus = C.AA_SPI_SLAVE_READ_ERROR
	errSpiSlaveTimeout       AardvarkStatus = C.AA_SPI_SLAVE_TIMEOUT
	errSpiDroppedExcessBytes AardvarkStatus = C.AA_SPI_DROPPED_EXCESS_BYTES

	// GPIO codes (-400 to -499)
	errGpioNotAvailable AardvarkStatus = C.AA_GPIO_NOT_AVAILABLE

	// I2C bus monitor codes (-500 to -599)
	errI2cMonitorNotAvailable AardvarkStatus = C.AA_I2C_MONITOR_NOT_AVAILABLE
	errI2cMonitorNotEnabled   AardvarkStatus = C.AA_I2C_MONITOR_NOT_ENABLED
)

var (
	ErrNil                    error = nil
	ErrUnableToLoadLibrary          = errors.New("Unable To Load Library")
	ErrUnableToLoadDriver           = errors.New("Unable To Load Driver")
	ErrUnableToLoadFunction         = errors.New("Unable To Load Function")
	ErrIncompatibleLibrary          = errors.New("Incompatible Library")
	ErrIncompatibleDevice           = errors.New("Incompatible Device")
	ErrCommunicationError           = errors.New("Communication Error")
	ErrUnableToOpen                 = errors.New("Unable To Open")
	ErrUnableToClose                = errors.New("Unable To Close")
	ErrInvalidHandle                = errors.New("Invalid Handle")
	ErrConfigError                  = errors.New("Config Error")
	ErrI2cNotAvailable              = errors.New("I2C Not Available")
	ErrI2cNotEnabled                = errors.New("I2C Not Enabled")
	ErrI2cReadError                 = errors.New("I2C Read Error")
	ErrI2cWriteError                = errors.New("I2C Write Error")
	ErrI2cSlaveBadConfig            = errors.New("I2C Slave Bad Config")
	ErrI2cSlaveReadError            = errors.New("I2C Slave Read Error")
	ErrI2cSlaveTimeout              = errors.New("I2C Slave Timeout")
	ErrI2cDroppedExcessBytes        = errors.New("I2C Dropped Excess Bytes")
	ErrI2cBusAlreadyFree            = errors.New("I2C Bus Already Free")
	ErrSpiNotAvailable              = errors.New("SPI Not Available")
	ErrSpiNotEnabled                = errors.New("SPI Not Enabled")
	ErrSpiWriteError                = errors.New("SPI Write Error")
	ErrSpiSlaveReadError            = errors.New("SPI Slave Read Error")
	ErrSpiSlaveTimeout              = errors.New("SPI Slave Timeout")
	ErrSpiDroppedExcessBytes        = errors.New("SPI Dropped Excess Bytes")
	ErrGpioNotAvailable             = errors.New("GPIO Not Available")
	ErrI2cMonitorNotAvailable       = errors.New("I2C Monitor Not Available")
	ErrI2cMonitorNotEnabled         = errors.New("I2C Monitor Not Enabled")
)

var errors_ = map[AardvarkStatus]error{
	errNil:                    ErrNil,
	errUnableToLoadLibrary:    ErrUnableToLoadLibrary,
	errUnableToLoadDriver:     ErrUnableToLoadDriver,
	errUnableToLoadFunction:   ErrUnableToLoadFunction,
	errIncompatibleLibrary:    ErrIncompatibleLibrary,
	errIncompatibleDevice:     ErrIncompatibleDevice,
	errCommunicationError:     ErrCommunicationError,
	errUnableToOpen:           ErrUnableToOpen,
	errUnableToClose:          ErrUnableToClose,
	errInvalidHandle:          ErrInvalidHandle,
	errConfigError:            ErrConfigError,
	errI2cNotAvailable:        ErrI2cNotAvailable,
	errI2cNotEnabled:          ErrI2cNotEnabled,
	errI2cReadError:           ErrI2cReadError,
	errI2cWriteError:          ErrI2cWriteError,
	errI2cSlaveBadConfig:      ErrI2cSlaveBadConfig,
	errI2cSlaveReadError:      ErrI2cSlaveReadError,
	errI2cSlaveTimeout:        ErrI2cSlaveTimeout,
	errI2cDroppedExcessBytes:  ErrI2cDroppedExcessBytes,
	errI2cBusAlreadyFree:      ErrI2cBusAlreadyFree,
	errSpiNotAvailable:        ErrSpiNotAvailable,
	errSpiNotEnabled:          ErrSpiNotEnabled,
	errSpiWriteError:          ErrSpiWriteError,
	errSpiSlaveReadError:      ErrSpiSlaveReadError,
	errSpiSlaveTimeout:        ErrSpiSlaveTimeout,
	errSpiDroppedExcessBytes:  ErrSpiDroppedExcessBytes,
	errGpioNotAvailable:       ErrGpioNotAvailable,
	errI2cMonitorNotAvailable: ErrI2cMonitorNotAvailable,
	errI2cMonitorNotEnabled:   ErrI2cMonitorNotEnabled,
}

func getError(status C.int) error {
	if e, ok := errors_[AardvarkStatus(status)]; ok {
		return e
	}
	return errors.New("unmapped error")
}

// PortNotFree
const PortNotFree = C.AA_PORT_NOT_FREE

// FindDevices returns a list of ports with Aardvark devices attached.
// Devices that are in use are ORed with PortNotFree (0x8000)
func FindDevices() (devices []uint16) {
	devs := make([]uint16, 256)
	cnt := int(C.c_aa_find_devices(256, (*C.u16)(unsafe.Pointer(&devs[0]))))
	if cnt == 0 {
		return
	}
	// XXX: should this be a for loop?
	return append(devices, devs...)
}

// GetUnusedDevices scans a device list returning an unused device list.
func GetUnusedDevices(devices []uint16) (unused []uint16) {
	for _, v := range devices {
		if (v & PortNotFree) == 0 {
			continue
		}
		unused = append(unused, v&^0x8000)
	}
	return
}

// DevInfo holds device port and id pair.
type DevInfo struct {
	Port uint16
	ID   uint32
}

// FindDevicesExt returns DevInfo pairs.
func FindDevicesExt() (deviceInfo []DevInfo) {
	devices := make([]uint16, 256)
	ids := make([]uint32, 256)
	cnt := int(C.c_aa_find_devices_ext(256, (*C.u16)(unsafe.Pointer(&devices[0])),
		256, (*C.u32)(unsafe.Pointer(&ids[0]))))
	if cnt == 0 {
		return
	}

	for j, val := range devices {
		deviceInfo = append(deviceInfo, DevInfo{val, ids[j]})
	}
	return
}

// Open attempts to open the device at the designated port.
func Open(port int) (*Device, error) {
	aa := C.c_aa_open(C.int(port))
	if int(aa) == 0 {
		return nil, ErrUnableToOpen
	}
	return &Device{aa: aa}, nil
}

type AardvarkVersion = C.AardvarkVersion
// type AardvarkVersion struct {
// 	Software      uint16
// 	Firmware      uint16
// 	Hardware      uint16
// 	Sw_req_by_fw  uint16 // Firmware requires that software must be >= this version.
// 	Fw_req_by_sw  uint16 // Software requires that firmware must be >= this version.
// 	Api_req_by_sw uint16 // Software requires that the API interface must be >= this version.
// }

type AardvarkExt = C.AardvarkExt
// type AardvarkExt struct {
// 	Version  AardvarkVersion // Version matrix
// 	Features int32           // Features of this device
// }

// OpenExt attempts to open the device at port.
// Returns a Device object as well as the device's extended
// info.
// func OpenExt(port int) (*Device, *AardvarkExt, error) {
// 	extInfo := &AardvarkExt{}
// 	aa := C.c_aa_open_ext(C.int(port), (*C.struct_AardvarkExt)(extInfo))
// 	if aa == nil {
// 		return nil, nil, ErrUnableToOpen
// 	}
// 	return &Device{AardvarkExt: aa}, extInfo, nil
// }

// Close closes the device.
func (d *Device) Close() error {
	return getError(C.c_aa_close(d.aa))
}

// Port returns the device's port number.
func (d *Device) Port() (int, error) {
	p := C.c_aa_port(d.aa)
	if p < 0 {
		return -1, getError(p)
	}
	return int(p), nil
}

const (
	featureSPI        = C.AA_FEATURE_SPI
	featureI2C        = C.AA_FEATURE_I2C
	featureGPIO       = C.AA_FEATURE_GPIO
	featureI2CMonitor = C.AA_FEATURE_I2C_MONITOR
)

// Features returns the device's feature set.
func (d *Device) Features() (spi, i2c, gpio, i2cMonitor bool, err error) {
	f := C.c_aa_features(d.aa)
	if f < 0 {
		err = getError(f)
		return
	}
	if (f & featureSPI) == featureSPI {
		spi = true
	}
	if (f & featureI2C) == featureI2C {
		i2c = true
	}
	if (f & featureGPIO) == featureGPIO {
		gpio = true
	}
	if (f & featureI2CMonitor) == featureI2CMonitor {
		i2cMonitor = true
	}
	return
}

// UniqueID returns the device's unique ID.
func (d *Device) UniqueID() uint32 {
	return uint32(C.c_aa_unique_id(d.aa))
}

type LogIO = C.int
var (
	LogStdout LogIO = LogIO(C.AA_LOG_STDOUT)
	LogStderr LogIO = LogIO(C.AA_LOG_STDERR)
)

// Log logs either stderr or stdout to file.
// func (d *Device) Log(stdio LogIO, file string) error {
// 	if d.fp != nil {
// 		return fmt.Errorf("the file %s is open, only one log file per session", file)
// 	}
// 	f := C.CString(file)
// 	defer C.free(unsafe.Pointer(f))

// 	m := C.CString("w")
// 	defer C.free(unsafe.Pointer(m))


// 	fp := C.open(f, C.int(C.O_WRONLY | C.O_CREAT | C.O_TRUNC))
// 	if fp == -1 {
// 		return fmt.Errorf("unable to open %s", file)
// 	}
// 	d.fp = fp

// 	i := C.c_aa_log(d.aa, C.int(stdio), fp)
// 	if i < 1 {
// 		C.close(fp)
// 	}
// 	return getError(i)
// }

// Version returns the device's version matrix.
func (d *Device) Version() (v *AardvarkVersion, err error) {
	v = &AardvarkVersion{}
	err = getError(C.c_aa_version(d.aa, v))
	return
}

// AardvarkConfig
type AardvarkConfig = C.AardvarkConfig
var (
	ConfigGpioOnly AardvarkConfig = C.AA_CONFIG_GPIO_ONLY
	ConfigSpiGpio  AardvarkConfig = C.AA_CONFIG_SPI_GPIO
	ConfigGpioI2c  AardvarkConfig = C.AA_CONFIG_GPIO_I2C
	ConfigSpiI2c   AardvarkConfig = C.AA_CONFIG_SPI_I2C
	ConfigQuery    AardvarkConfig = C.AA_CONFIG_QUERY
)

// Config configures the device by enabling/disabling
// I2C, SPI, and GPIO functionality.
func (d *Device) Config(config AardvarkConfig) error {
	return getError(C.c_aa_configure(d.aa, config))
}

// TargetPower
type TargetPower = C.u08

var (
	TargetPowerNone  TargetPower = C.AA_TARGET_POWER_NONE
	TargetPowerBoth  TargetPower = C.AA_TARGET_POWER_BOTH
	TargetPowerQuery TargetPower = C.AA_TARGET_POWER_QUERY
)

// TargetPower configures the device's target power pins.
func (d *Device) TargetPower(powerMask TargetPower) error {
	return getError(C.c_aa_target_power(d.aa, powerMask))
}

//
// SleepMs sleeps for the specified number of milliseconds.
// Returns the number of milliseconds slept.
func (d *Device) SleepMs(ms uint32) uint32 {
	return uint32(C.c_aa_sleep_ms(C.u32(ms)))
}

// AsyncPoll
type AsyncPoll C.int

var (
	AsyncNoData     AsyncPoll = C.AA_ASYNC_NO_DATA
	AsyncI2cRead    AsyncPoll = C.AA_ASYNC_I2C_READ
	AsyncI2cWrite   AsyncPoll = C.AA_ASYNC_I2C_WRITE
	AsyncSpi        AsyncPoll = C.AA_ASYNC_SPI
	AsyncI2cMonitor AsyncPoll = C.AA_ASYNC_I2C_MONITOR
)

// AsyncPoll checks if there are any asynchronous messages pending
// for processing. The call will time out for the number of milliseconds
// defined, if timeout is < 0 the call will block until data is received,
// if timeout is == 0 the call returns immediately.
func (d *Device) AsyncPoll(timeout int) AsyncPoll {
	return AsyncPoll(C.c_aa_async_poll(d.aa, C.int(timeout)))
}

// SpiBitRate sets the SPI bit rate in kilohertz.  If a zero is passed as the
// bitrate, the bitrate is unchanged and the current bitrate is returned.
func (d *Device) SpiBitRate(bitrateKHz int) int {
	return int(C.c_aa_spi_bitrate(d.aa, C.int(bitrateKHz)))
}

// SpiPolarity
type SpiPolarity = C.AardvarkSpiPolarity

var (
	SpiPollRisingFalling SpiPolarity = C.AA_SPI_POL_RISING_FALLING
	SpiPolFallingRising  SpiPolarity = C.AA_SPI_POL_FALLING_RISING
)

// SpiPhase
type SpiPhase = C.AardvarkSpiPhase

var (
	SpiPhaseSampleSetup SpiPhase = C.AA_SPI_PHASE_SAMPLE_SETUP
	SpiPhaseSetupSample SpiPhase = C.AA_SPI_PHASE_SETUP_SAMPLE
)

// SpiBitOrder
type SpiBitOrder = C.AardvarkSpiBitorder

var (
	SpiBitOrderMSB = C.AA_SPI_BITORDER_MSB
	SpiBitOrderLSB = C.AA_SPI_BITORDER_LSB
)

// SpiConfig configures the SPI master or slave interface.
func (d *Device) SpiConfig(polarity SpiPolarity, phase SpiPhase, bitorder SpiBitOrder) error {
	return getError(C.c_aa_spi_configure(d.aa, polarity, phase, bitorder))
}

// SpiWrite writes a stream of bytes to the downstream SPI slave device.
func (d *Device) SpiWrite(outBuf []byte, inByteCnt int) ([]byte, error) {
	inBuf := make([]byte, inByteCnt)
	err := getError(C.c_aa_spi_write(d.aa,
		C.u16(len(outBuf)),
		(*C.u08)(unsafe.Pointer(&outBuf[0])),
		C.u16(inByteCnt),
		(*C.u08)(unsafe.Pointer(&inBuf[0]))))
	return inBuf, err
}

// SpiSlaveEnable enables the device as a SPI slave.
func (d *Device) SpiSlaveEnable() error {
	return getError(C.c_aa_spi_slave_enable(d.aa))
}

// SpiSlaveDisable disables the device as a SPI slave.
func (d *Device) SpiSlaveDisable() error {
	return getError(C.c_aa_spi_slave_disable(d.aa))
}

// SpiSlaveSetResponsee set the response to the master
// when in slave mode.
func (d *Device) SpiSlaveSetResponsee(response string) error {
	s := C.CString(response)
	defer C.free(unsafe.Pointer(s))
	return getError(C.c_aa_spi_slave_set_response(d.aa, C.u08(len(response)), (*C.u08)(unsafe.Pointer(s))))
}

// SpiSlaveRead reads byteCnt bytes from the SPI slave line.
func (d *Device) SpiSlaveRead(byteCnt int) ([]byte, error) {
	buf := make([]byte, byteCnt)
	err := getError(C.c_aa_spi_slave_disable(d.aa))
	return buf, err
}

// SpiSSPolarity
type SpiSSPolarity = C.AardvarkSpiSSPolarity

var (
	SpiSSActiveLow  SpiSSPolarity = C.AA_SPI_SS_ACTIVE_LOW
	SpiSSActiveHigh SpiSSPolarity = C.AA_SPI_SS_ACTIVE_HIGH
)

// MasterSSPolarity changes the output polarity on the SS line.
// Note: When configured as an SPI slave, the device will always
// be setup with SS as active low. Hence this function
// only affects the SPI master functions on the device.
func (d *Device) MasterSSPolarity(polarity SpiSSPolarity) error {
	return getError(C.c_aa_spi_master_ss_polarity(d.aa, polarity))
}
