# Handling Deployment Order <img src="https://tutorialedge.net/images/golang.png"  width="100" align="right" />

 **Handling Deployment Order is a tool to deploy different components (services), each of which requires on some other components.**

 **You can add YAML or JSON file to be deploy and give back components sequence.**
 
 
<img src="https://cdn-images-1.medium.com/max/1200/0*ABKj5V2gw7tezUc5." width="250">

## Package requires

To get the package, execute:
```sh
go get gopkg.in/yaml.v2
```
To import this package, add the following line to your code:
```sh
import "gopkg.in/yaml.v2"
```
For more details, see [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2).

## How To Use

#### In package:

Call function **YamlFileReader(path string)Components** then **DeployOrder(components Components)Components** in main function.

````
var components = repository.YamlFileReader("resource/YamlFileReader_Test.yaml")
var orderComponents = service.DeployOrder(components)
````
## Examples

Yaml file:

````
components:
  - component: A
    requires:
      - D
      - E
  - component: C
    requires:
      - B
  - component: D
    requires:
      - B
  - component: E
    requires:
  - component: F
    requires:
  - component: B
    requires:
````

Main file:

````
var components = repository.YamlFileReader("YamlFile.yaml")
var orderComponents = service.DeployOrder(components)
	
for index, component := range orderComponents {
	fmt.Println(index, ":", component.Name)
}
````

Output:

````
0 : B
1 : D
2 : E
3 : A
4 : C
5 : F
````

#### With HTTP request:

Web server which handles an http request via rest api.

Listen on 8081 port.

It take JSON format as format below:

```json
{
  "components": [
    {
      "component": "A",
      "requires": [
        "D",
        "E"
      ]
    },
    {
      "component": "C",
      "requires": [
        "B"
      ]
    },
    {
      "component": "D",
      "requires": [
        "B"
      ]
    },
    {
      "component": "E",
      "requires": null
    },
    {
      "component": "F",
      "requires": null
    },
    {
      "component": "B",
      "requires": null
    }
  ]
}
```
#### Console Client which sends requests to a web server:


```
Choose type of read data. NOTE: Pick up with number.
1. Read from file
2. Read from console
3. Auto generated YAML files with Components
```

Choose 1
```
Give absolute path of file:
    Here you can give path.
```

Choose 2
```
Тhe program will stop to read after you are writing an END!
{
  "components": [
    {
      "component": "A",
      "requires": [
        "D",
        "E"
      ]
    },
    {
      "component": "C",
      "requires": [
        "B"
      ]
    },
    {
      "component": "D",
      "requires": [
        "B"
      ]
    },
    {
      "component": "E",
      "requires": null
    },
    {
      "component": "F",
      "requires": null
    },
    {
      "component": "B",
      "requires": null
    }
  ]
}
end
```

Choose 3
```
How many files to generated ?
 Here you should give number of generated YAML files, you want it.
```

#### Slack Message Sender  
###### gives result of determining components

```
SendMessageInSlackChannel(url string, message string)
```
