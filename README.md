# AI Interactor

AI Interactor is a web server that allows interaction with multiple AI models by sending prompts.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine or edge device for development and testing purposes.

### Prerequisites

Install dependencies

```````
make install-deps
```````

#### Building the application on edge devices
Compile the application

````
make build/edge
```````

Download the applicaiton to your edge device (take into account that SSH access is needed)

````
make deploy USER=<username on the edge device> DEVICE_URL=<the device url>
````

#### Building the application on a regular machine
Compile the application

````
make build
```````

Run the app

````
ai-iteractor
````

