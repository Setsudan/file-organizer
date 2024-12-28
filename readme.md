# File Organizer: Automated Downloads Sorting

Welcome to the **File Organizer** project! This **Go** script automatically organizes the contents of your **Downloads** folder into predefined categories (Documents, Videos, Audio, Pictures, etc.). It moves any files it recognizes into a “Downloaded ___” folder within the appropriate directory, all based on file extensions.

---

## Features

1. **Future-Proof File Paths**:  
   Automatically locates a user’s home directory and derives all needed paths (Documents, Videos, Music, Pictures, Others) without any hardcoding.

2. **Comprehensive File Categories**:  
   Groups files by extensions for Documents, Videos, Audio, and Pictures. Customize it by adding or removing categories in the `fileCategories` map.

3. **Car-Themed Easter Egg**:  
   Keep an eye out for the fun tribute to endurance racing and cars!

4. **Developer-Friendly**:  
   Code is structured to mix simple OOP practices with functional helpers, ensuring clarity and maintainability.

5. **Startup Execution**:  
   Optionally run the program at system startup to keep your downloads sorted automatically.

---

## Getting Started

1. **Prerequisites**  
   - [Go](https://golang.org/) (1.18 or higher recommended)
   - A computer running Windows, macOS, or Linux

2. **Download or Clone This Repo**  

   ```bash
   git clone https://github.com/setsudan/file-organizer-go.git
   cd file-organizer-go
   ```

3. **Build the Executable**  

   ```bash
   go build -o file-organizer
   ```

   - On Windows, you may want to hide the console window:

     ```bash
     go build -ldflags="-H=windowsgui" -o file-organizer.exe
     ```

4. **Run the Script**  

   ```bash
   ./file-organizer
   ```

   The script will:
   1. Look in your **Downloads** folder
   2. Identify each file’s category
   3. Move files to a `Downloaded [Category]` folder inside **Documents**, **Videos**, **Music**, **Pictures**, or **Others**.

---

## Run on Startup (Windows Example)

1. **Generate Executable** (see above).
2. **Open Your Startup Folder**  
   Press **Win + R**, type `shell:startup`, and hit Enter.
3. **Place or Link** the compiled `.exe` in this folder.  
   Your file organizer will now run automatically every time you log in, quietly sorting your downloads.

---

## Customizing

- **File Categories**  
  Open the script and locate the `fileCategories` map. Add or remove extensions as you see fit.
- **Directory Paths**  
  If you want to change the target folders, edit the code where it assigns `docDir`, `videoDir`, etc. But by default, it uses system directories in your home folder.
- **Logging / Feedback**  
  Feel free to adjust or remove the console output for a more silent experience.

---

## Contributing

Feel free to fork this repository and submit pull requests with improvements:

- Adding new categories
- Improving file extension coverage
- Enhancing code clarity
- Expanding OS compatibility

---

## License

This project is distributed under the **MIT License**. Please see the [LICENSE](LICENSE) file for details.

---
