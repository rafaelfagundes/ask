#!/bin/bash

# Function to wait for keypress
press_key() {
    echo -e "\n$1"
    read -n 1 -s
    echo
}

echo -e "\n\n=== Testing Config Directory ==="
press_key "Press any key to show the config directory 📁"
ask -c

echo -e "\n\n=== Testing Help ==="
press_key "Press any key to display help menu ℹ️"
ask --help

echo -e "\n\n=== Testing Main Command ==="
press_key "Press any key to test a direct question about Go 🤔"
ask "What is Go?"

echo -e "\n\n=== Testing No Pager Option ==="
press_key "Press any key to test no-pager option 📑"
ask --no-pager "Print a short response"

echo -e "\n\n=== Testing History Commands ==="
press_key "Press any key to list all history entries 📜"
ask history

press_key "Press any key to show history entry #1 🔍"
ask history 1

echo -e "\n\n=== Testing Last Command ==="
press_key "Press any key to show last response ⏮️"
ask last

press_key "Press any key to show last response without pager 📄"
ask last --no-pager

echo -e "\n\n=== Testing History Delete ==="
press_key "Press any key to delete history entry #1 🗑️"
ask history delete 1

press_key "Press any key to delete all history (requires confirmation) ⚠️"
ask history delete all

echo -e "\n\n=== Testing History Clear ==="
press_key "Press any key to show final history state 📊"
ask history

echo -e "\n\n=== Unit Tests ==="
press_key "Press any key to run unit tests 🧪"
go test ./...

echo -e "\n\n>>>> Done! 🎉 <<<<"