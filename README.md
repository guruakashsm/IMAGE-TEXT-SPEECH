<!DOCTYPE html>
<html>

<head>
    <title>IMAGE-TEXT-SPEECH README</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }

        h1 {
            color: #337ab7;
        }

        h2 {
            color: #333;
        }

        p {
            margin: 10px 0;
        }

        code {
            background-color: #f8f8f8;
            padding: 2px 6px;
            border: 1px solid #ddd;
        }
    </style>
</head>

<body>
    <h1>IMAGE-TEXT-SPEECH</h1>
    <p>This project allows you to convert English text from an image to Tamil, English, and Hindi text and speech using ML techniques.</p>

    <h2>Installation and Run</h2>
    <ol>
        <li>Install Go and Python.</li>
        <li>Install PIP for Python.</li>
        <li>Run the following commands to install required Python packages:</li>
    </ol>
    <pre><code>
pip install pytesseract
pip install Pillow
pip install googletrans==4.0.0-rc1
pip install gTTS
pip install Flask
    </code></pre>
    <ol start="4">
        <li>Go to the project directory and run:</li>
    </ol>
    <pre><code>
go mod tidy
    </code></pre>
    <ol start="6">
        <li>Run the Python and Go applications using the following command:</li>
    </ol>
    <pre><code>
go run main.go
    </code></pre>
    <ol start="7">
        <li>Access the application locally and try IMAGE &rarr; TEXT &rarr; SPEECH conversion.</li>
    </ol>

    <h2>Features</h2>
    <ul>
        <li>Complete BACKEND AND FRONTEND for image-to-text and text-to-speech conversion.</li>
        <li>Supports 3 languages: Tamil, English, and Hindi.</li>
        <li>Improved accuracy using ML techniques.</li>
    </ul>
</body>

</html>
