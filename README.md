This is an event driven framework I created to complete my understanding of concurrent programming in Go. User defines Services by implementing Init(), Run(), and helper methods.

## Framework

-**Main.go**: Instantiates and runs the Framework.

-**Framework**: An event loop, listens for and distributes events to Services.

-**Framework_Definitions**: Contains definitions for the Event type and for the Service type. Also contains methods that are common to all Services.

## Services

-**Services**: Each app will have a set of Services in the same package. The package should contain a Services.go file which contains custom Event/type definitions and individual files specific to each Service.

-**Examples**: As of this writing there are a few example implementations of apps using the Framework. GIS_Services is a command line tool that lets the user choose data and ranges and maps it onto an image of the world. Use GIS_Services as a template. IO_Service and Menu_Service are ready to go services for getting command line input, output, and menu navigation.

More details are available in the Services_Example folder.
