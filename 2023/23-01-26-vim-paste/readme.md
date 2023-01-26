# set vim paste

For any operating system and client

Regardless of which operating system or client you use to connect to a remote host and use Vim, you can use regular copy/paste to copy text while in a Vim session. Follow this process:

* Use Vim to open the file that you want to edit.
* Copy some text by highlighting it with your mouse and typing Cmd-C or Ctr-C.
* In Vim, go into insert mode by typing “i” (notice that “– INSERT –” appears on the bottom of the Vim session).
* Type Cmd-V or Ctr-V to paste the text.
If you are copying text for which the formatting should not change, beware that the above process can introduce changes to the formatting. For example, if you are copying Python code into a file using Vim, be prepared for the indentation to be altered, and therefore your program to not work as expected. In Python, code blocks like loops are denoted using text indentation.

To avoid this from happening, you can use Vim’s paste mode. When you enable paste mode, Vim will not auto-indent any text that you paste. To enable paste mode, follow this process:

* In Vim, ensure you are command mode by hitting the Esc key.
* Type “:set paste” to enter command mode.
* Type “i” to enter paste mode. Notice the “– INSERT (paste) –” at the bottom of the Vim window.
* Paste your Python code into Vim. Indentation should be as in the original
* To exit paste mode, type “:set nopaste”. Notice the “– INSERT –” at the bottom of the Vim window. This means you are back to normal insert mode.
