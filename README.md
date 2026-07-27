## Mana Royale
A fun blend of AI and bluffs. You will be competing against a sly Devil's advocate.
He never hesitates to mislead into despair. He guards a valueable treasure. So he is entrusted to deceive any one who comes nearby.
You, the player, will be given 3 cards each turn. You have guess the cards face down and choose. The advocate will give you hints but are they true? Or its your source of despair.

The Advocate persona is represented by a local LLM. He has a small window to be truthful or is it a reverse psychological manipulation. Find out and defeat the liar.

### Card types:
1. Damage : -2xDMG, -1xDMG, 0xDMG, 1xDMG, 2xDMG ( negative number means multiplier reduce damage to advocate)
2. Projectile-Spell : 25, 50, 100, 200, 500 (Direct damage to Devil's advocate)
3. Epic-Spell : Heal, Poison, Sleep, Hex, Instant-Death (Spells)

## Epic Spells:
* **Hex** the effect of **target switching**. Meaning if you are under Hex, and you use 2xDMG, then instead of the devil the player will be hit. It effects **all** other spells
* Instant-Death puts the devil into hell. If hex is active then player dies
(More polished rules will arrive at future. Idea remains the same)

## Tech-stack: 
1. Go
2. Ollama

## Set up:

1. Note: If binary not provided/published, use Go to build and run the project.
2. Use Git to load the project -
   ```bash
   git clone https://github.com/Soumen-No-Hako/Mana-Royale.git
   ```
   or
   ```bash
   git clone git@github.com:Soumen-No-Hako/Mana-Royale.git
   ```
4. Check **if Ollama exists** in your device my running ``` ollama --version ``` in command line. Or by searching "ollama"
   You can install by single command
   Ollama simple installation command -
   ```sh
   curl -fsSL https://ollama.com/install.sh | sh
   ```
6. The project by default requires a **llama3.2 LLM weights**, install it using
   ```bash
   ollama pull llama3.2
   ```
7. (Optional) Ollama server automatically runs. If not run by ``` ollama serve ```
8. Now to run Mana-Royale
   ```bash
   go run mana-royale.go
   ```
   If executed correctly the output should look like -
   <img width="1920" height="715" alt="image" src="https://github.com/user-attachments/assets/3808fd00-d508-4203-8b71-1c276c7fdacd" />
9. ### **Play by entering 1, 2 or 3 as std input.**

## Modifying devil-modelfile
 the **devil-modelfile** is located at [Templates/devil-modelfile](https://github.com/Soumen-No-Hako/Mana-Royale/blob/main/Templates/devil-modelfile)
 The game code has a **parser** that can parse the modelfile properties.

 Which means 
 1. feel free to **change** the modelfile parameter as needed by your device.
 2. Choose any **compatible gguf** file by adding it to "FROM " directive
    Example:
    ```zsh
    FROM gemma4
    PARAMETER num_ctx 2048
    ....
    ```
    or
    ```zsh
    FROM "~/path/to/gguf"
    PARAMETER temperature 0.75
    SYSTEM """
    you are a devil who ...........
    ......
    """

## Reference
1. Go reference - [install-go](https://go.dev/doc/install)
2. Ollama reference - [ollama/docs](https://docs.ollama.com/)
