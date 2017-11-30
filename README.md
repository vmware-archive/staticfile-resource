# staticfile-resource
A concourse resource to take arbitrary data passed in from a pipeline and writes
them to files to be used in a task.

## Source Configuration

* `files`: **Required.** An array of files to be written by the resource.

  * `filename`: **Required.** The name of the file to write to.

  * `data`: **Required.** The data that is written to the file.

## Behavior

### `in`: Write static files with data from the resource's source configuration.

Writes files to the destination directory with the arbitrary data from the
resource's source configuration.

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
    files:
    - filename: {{filename}}
      data: {{configuration_data}}
```
