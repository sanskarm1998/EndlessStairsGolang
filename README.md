# EndlessStairsGolang

Endless Stairs is a terminal-based game written in Go. Jump left or right to climb the endless stairs. The game gets faster as your score increases! Now with power-ups, collectibles, and improved UI.

## How to Play
- Use ←/→ arrow keys or l/r keys to jump to the next stair.
- The next stair is shown above your character.
- The game gets faster as your score increases.
- The game ends if you jump in the wrong direction or run out of time.
- Your score, coins, gems, and time left are always shown at the top of the screen.

## Run the Game
```sh
go run ./cmd/main.go
```

## Requirements
- Go 1.18 or newer
- Terminal that supports ANSI escape codes (for screen clearing)

## Stair Types

- **Normal Stair**: Standard stair. Jump in the correct direction to score a point. (60%)
- **Falling Stair**: The stair collapses if you land on it! You barely make it, but your score does not increase. The time to select is halved (but never below 0.9 seconds). (5%)
- **Spiked Stair**: Ouch! Landing on this stair reduces your score by 1 (but not below 0). The game continues. (5%)
- **Reverse Polarity Stair**: Controls are reversed for this stair! Left becomes right and right becomes left for this round. Watch for the warning message! (5%)
- **Super Stair**: Jumping on this stair gives you a big bonus! Score +5 points. Watch for the special ASCII art ([***]). (3%)
- **Slow Time Stair**: Slows down the timer for the next 5 stairs. (3%)
- **Double Points Stair**: Double points for the next 5 stairs. (2%)
- **Shield Stair**: Grants a shield that protects you from the next hazard (falling or spiked stair). (5%)
- **Coin Stair**: Collect a coin for your total. (5%)
- **Gem Stair**: Collect a rare gem for your total. (2%)

Each stair type is visually distinct in the ASCII art.

## Power-Ups and Collectibles

- **Slow Time**: Temporarily increases the time you have to make decisions.
- **Double Points**: Temporarily doubles the points you earn for each correct jump.
- **Shield**: Protects you from the next falling or spiked stair.
- **Coins & Gems**: Collect these for fun and future unlockables!

## UI Improvements
- Score, coins, gems, and time left are now displayed at the top of the screen for better visibility.
- Reduced vertical space for a more compact and engaging play area.

## Stair Generation Probabilities
- 60% Normal, 5% Falling, 5% Spiked, 5% Reverse, 3% Super, 3% Slow Time, 2% Double, 5% Shield, 5% Coin, 2% Gem


