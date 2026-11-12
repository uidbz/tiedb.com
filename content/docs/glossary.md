---
title: Glossary
type: docs
weight: 90
---

A shared vocabulary for the library, the apps and the agent. The same list is
available to the prose as a `term` shortcode.

## Library

{{< term "triple" >}}
{{< term "content hash" >}}
{{< term "filehost" >}}
{{< term "tiedir" >}}
{{< term "dir-type" >}}
{{< term "collection" >}}
{{< term "tag" >}}
{{< term "mount" >}}
{{< term "verify" >}}

## Agent

{{< term "harness" >}}
{{< term "overlay" >}}
{{< term "promote" >}}
{{< term "probe" >}}
{{< term "session" >}}
{{< term "Risor" >}}

## Everything else

| Term | Meaning |
|---|---|
| **blob** | The bytes of a file, stored once under their hash. |
| **reverse index** | An opt-in relation index that answers "who points at this value?". |
| **tombstone** | A deleted record whose slot is reused; the file does not grow while holes exist. |
| **file version** | A previous manifest/tree that contained the same content hash. |
| **dir-type** | The type of a directory: audio, image, video, document, archive, or a custom one. |
| **pwplay** | The PipeWire audio engine used by tie-audio. |
| **fynerisor** | Risor bindings for Fyne, used for scripted panes and tie-browser scenes. |
| **tiestyle** | The shared Fyne theme package behind the apps' surfaces and accents. |
