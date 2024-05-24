# German

## Description
A simple game to learn German words. The game will display the English definition of the word along with a sentence in English that uses the word. The player must then type the German word that corresponds to the English definition.

## How to play
1. Run the game
2. Enter the number of words you want to learn

## There are two modes
1. **Normal Mode**: The game will display the English definition of the word along with a sentence in English that uses the word. The player must then type the German word that corresponds to the English definition. The list of words is from the `./data/KnownWords.json` file.
2. **Challenge Mode**: The game will display the English definition of the word along with a sentence in English that uses the word. The player must then type the German word that corresponds to the English definition. The list of words is from the `./data/Top1000.json` file.

## How to add more words
1. Open the `./data/KnownWords.json` file
2. Add the words you want to learn in the following format:
```json
{
  "word": "ich",
  "EnglishSentence": "I like rain.",
  "pos": ["pronoun"],
  "definition": ["I"],
  "description": "First person singular pronoun. Used to refer to oneself."
}
```
3. Run the game

## How to remove words
1. Open the `./data/KnownWords.json` file
2. Remove the word you want to remove
3. Run the game

## How to update words
1. Open the `./data/KnownWords.json` file
2. Update the word you want to update
3. Run the game

## CLI Commands
- `--help | -h`: Display the help menu
- `--challenge | -h`: Play the game in challenge mode (default is normal mode)

## Command Examples
```bash
german --help OR german -h

german --challenge | german -c
```