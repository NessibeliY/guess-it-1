# guess-it
This is a program that given a number as standard input, prints out a range in which the next number provided should be.
This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

## Usage

You will first need copy the `student/` folder (provided by the student) in which you will see the student guessing program along with a file called `script.sh`. This file should be an executable shell script that runs the student program if you are in the root folder <a href="https://assets.01-edu.org/guess-it/guess-it.zip">`guess-it/`</a>. The filesystem should look somethings like this:

```console
─ guess-it/
├── ai/
│   ├── big-range
│   └── ...
├── index.html
├── index.js
├── go.mod
└── ...
└── student/
    ├── ...
    └── script.sh

```

In order to test the student program, these commands should be ran to have the dependencies needed and to start the webpage on the port 3000:

```console
npm install
node server.js
```

You will need to run also this command inside the `ai/` directory to make the programs executable:

```console
chmod +x *
```

After opening your browser of preference in the port [3000](http://localhost:3000/), if you try clicking on any of the `Test Data` buttons, you will notice that in the Dev Tool/ Console there is a message which tells you that you need another guesser besides the student.

Adding a guesser is simple. You just need to add in the URL a guesser, in other words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose which of the random data set to test. After that you can wait for the program to test all of the values (boooooring), or you can click `Quick` in order to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays clicking on the `Clean` button after each test.
