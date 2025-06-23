# Slideshower

Slideshower is a simple image slideshow application written in Go using the Ebiten game library.

## Compiling the Application

To compile the Slideshower application, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [https://golang.org/](https://golang.org/).
2. Clone or download this repository to your local machine.
3. Open a terminal and navigate to the project directory.
4. Run the following command to compile the application: `go build .`

This will create an executable named `slideshower` (or `slideshower.exe` on Windows) in the current directory.

## Running the Application

To run Slideshower:

1. Copy the `slideshower` executable to a directory containing the images you want to display in the slideshow.
2. Run the application by double-clicking on it
3. To quit, either press `ESC` or `q` to quit.

By default, the application will:

- Display each image for 5 seconds
- Run in windowed mode (not fullscreen)
- Use the crossfade transition effect

## Configuring the Application

You can customize the behavior of Slideshower by creating a `config.yml` file in the same directory as the executable. Here's how to configure the options:

1. Create a file named `config.yml` in the same directory as the Slideshower executable.
2. Add the following content to the file:

```yaml
full_screen: false
speed_in_seconds: 5
effect: "crossfade"
```

3. Modify the values as desired:

* full_screen: Set to true for fullscreen mode, false for windowed mode.
* speed_in_seconds: Set the number of seconds each image is displayed.
* effect: Choose the transition effect. Options include:
    * "fade"
    * "crossfade"
    * "slide-from-left"
    * "slide-from-right"
    * "alternating-slide"
    * "spiral-wipe"
    * "bubble-melt" _(currently broken. it crashes the application)_
    * "random" (selects a random effect for each transition)

Example config.yml for a fullscreen slideshow with 10-second displays and random transitions:

```yaml
full_screen: true
speed_in_seconds: 10
effect: "random"
```

Save the config.yml file and run the application. It will automatically use the specified settings.
Enjoy your slideshow!

This README provides clear instructions on how to compile the application, run it with default settings, and customize its behavior using the `config.yml` file. It covers all the main points you requested, including:

1. How to compile the app using Go
2. How to run the app by copying the executable to a folder with images
3. The default settings (5-second display, windowed mode, crossfade effect)
4. How to create and use a `config.yml` file to configure the options

You can save this content as `README.md` in your project's root directory. Feel free to adjust any details or add more information as needed.
