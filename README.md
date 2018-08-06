# WebSynth

## Requirements
* Portaudio

## Running the synth

```bash
$ go get ./...
$ go run main.go
```

*Note: if the application panics with a message related to an invalid number of channels, you can try adjusting the [devices](https://github.com/dalloriam/websynth/blob/master/app/audio/backend.go#L21) used by portaudio. I will implement proper device detection with the next releases.*


## Using the synth

The synthesizer is controllable via a GraphQL API running on `http://localhost:8080/synth`.

### Display Mixer Status

```graphql
query {
    Mixer {
        Channels {
            Pan {
                Value
            }
            Volume {
                Value
            }
        }
    }
}
```

### Adjust a mixer channel's volume and pan

```graphql
mutation {
    Mixer {
        Channel(Idx: 0) {
            Volume {
                Set(Value: 1)
            }
            Pan {
                Set(Value: -1)
            }
        }
    }
}
```

### List Available Modules

*Note: For now, WebSynth only supports the oscillator module. This will obviously improve in the next releases.*

```graphql
query {
    Modules {
        List {
            ... on Oscillator {
                Frequency {
                    Value
                }
                Volume {
                    Value
                }
            }	
        }
    }
}
```

### Create an Oscillator

```graphql
mutation {
    Modules {
        Create {
            Oscillator 
        }
    }
}
```


### Attach a module to a mixer channel

```graphql
mutation {
    Mixer {
        Channel(Idx: 0) {
            Input {
                Attach(ModuleIdx: 0, ModuleField: "Sine") 
            }
        }
    }
}
```

### Detach a module from a mixer channel

```graphql
mutation {
    Mixer {
        Channel(Idx: 0) {
            Input {
                Detach
            }
        }
    }
}
``` 