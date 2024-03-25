# Raspberry pi `ddcutil`

> `ddcutil` is used to control monitor/screen settings

These steps were successful on a pi 3 running Raspbian lite 64bit OS.

## Install

```sudo apt install ddcutil```

## Error

Running any `ddcutil` command gave this error...

```
No /dev/i2c devices exist.
ddcutil requires module i2c-dev.
```

### Fix

```
# load module
sudo modprobe i2c-dev

# confirm it loaded
lsmod | grep i2c_dev

# enable module at boot
sudo echo "i2c-dev" >> /etc/modules
```

Then make sure I2C interface is enabled in Pi's configuration. `sudo raspi-config` and enable I2C in "Interfacing Options".

```sudo reboot```

## Usage

Detect active I2C bus

```
ddcutil detect
```

Get brightness of monitor (10 refers to brightness)

```
ddcutil getvcp 10
```

Set brightness of monitor (10 refers to brightness)

```
ddcutil setvcp 10 1  # low
ddcutil setvcp 10 50 # mid
ddcutil setvcp 10 99 # high
```

