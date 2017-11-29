# staticfile-resource
A concourse resource to take arbitrary data passed in from a pipeline and writes
it to a file to be used in a task.

## Source Configuration

* `filename`: *Required.* The name of the file to write to.

* `data`: *Required.* The data that is written to the file.

## Behavior

### `in`: Write a static file with source data.

Writes a file to the destination with the arbitrary data from the resource's
source configuration.

#### Parameters

* `file`: *Required.* Path to the file to upload, provided by an output of a task.

## Example Configuration

An example pipeline exists in the `example` directory.

### Resource

```
resource_types:
- name: staticfile
  type: docker-image
  source:
    repository: christianang/staticfile-resource

resources:
- name: configuration
  type: staticfile
  source:
    filename: {{filename}}
    data: {{configuration_data}}
```
