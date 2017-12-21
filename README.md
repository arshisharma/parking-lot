### Automated Parking lot
Application can accept command file
as well as commands in interactive mode.

#### Application build requirements
    go1.8 is required binaries. Executable permission enabled for
    `parking_lot`. Please make sure its in executable mode in your system
    Run
    `chmod 755 parking_lot`
    in command prompt if not.
    Please make sure your  ***GOPATH*** is set to application directory
    or set as folows.

   ***cd [solution directory]***

   ***export GOPATH=$(pwd)***

    Run above commands in your shell

#### Application has been built with following assumptions
1. Customers are nice enough to always park in the slots allocated to them.
2. Supports parking lot management of only one type of vehicle.
3. One Vehicle can occupy at most parking slot.
4. Only slot allocation and freeing is supported by this system.
5. This system doesn't manage the parking lot ticketing pricings.
6. All parking slots are of similar type. There are no categories of parking slot eg: slot with higher price or lower price.
7. There is only one entry/exit point for the parking lot.
