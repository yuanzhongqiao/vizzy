<div align="center">
    <img style="max-height: 300px" src="/app/src/assets/logo-with-text-red.png" alt="Vizzy smiling robot logo">
    <br>
    <hr>
    <img alt="Accuracy 77.3%" src="https://img.shields.io/badge/Accuracy-77.3%25-%23facc15">
    <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/rbren/vizzy/go-test">
    <img alt="MIT License" src="https://img.shields.io/badge/License-MIT-green">
</div>

# Vizzy
Vizzy uses ChatGPT to visualize any kind of data.

[Play with the demo](https://vizzy.rbren.io), or check out these sample projects:
* [Scripts from TV's "The Office"](https://vizzy.rbren.io/projects/393d0b17-bf2f-4476-81bd-82133e15d169)
* [Same-sex Marriage Over Time](https://vizzy.rbren.io/projects/0280e3b8-113d-43c0-a794-d022719a2833)
* [U.S. State Obesity Rates](https://vizzy.rbren.io/projects/b68f563c-90ee-40fb-b501-d4758764a331)
* [Cause of Death News Coverage](https://vizzy.rbren.io/projects/f7dafd12-6f00-4a8c-8c4c-7a4f51dc34f2)
* [U.S. National Parks](https://vizzy.rbren.io/projects/e09dcbe7-cf8b-4049-a642-e1935c80318c)

## About
A typical workflow goes like this:
* _**Upload**_ your data
  * Vizzy tells you what it can infer about the structure and meaning of the data
  * Example: "This is a CSV file containing facts about different car models, with the fields `year`, `make`, `model`, `horsepower` and `weight`
* _**Make any clarifications**_ or fixes to Vizzy's summary
* Tell Vizzy how you'd like to **visualize the data**
  * Vizzy will also give you some suggestions you can choose from
  * Example: "Draw a scatterplot of horsepower versus weight"
* _**Check out the results**_
  * More often than not, Vizzy gets it right!
  * But sometimes Vizzy may reply with an error message or a broken graph
  * See [Failure Patterns](FailurePatterns.md) for common issues and resolutions, or to report a new issue
* Tell Vizzy to _**make adjustments**_ or add new features
  * You can always revert back to a previous version if you're not happy with the change
  * Example: "Color the data points by manufacturer" or "Add a tooltip showing the make and model"
* Or _**edit the code yourself**_
  * This is faster and cheaper for trivial changes, like changing a color
  * It can also be helpful for issues Vizzy seems unable to address via prompts
* _**Repeat!**_

## Development
See [Contributing.md](Contributing.md) for instructions on running Vizzy locally.