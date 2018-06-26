# Validator PvP Scoreboard

### Requirements:

- List of Validator nodes in order of amount of stake
- Possible: Also include an uptime number
- Others?

### Potential Routes:

There are three potential routes here:
  - Talk with `figment.network` about using their's potentially with a different skin.
  - A scoreboard that polls the `/stake/validators` endpoint and forwards that data with some enrichment from
  - A scoreboard that parses through the whole chain, or just starts with the chain and stores the validators, along with some data for every block, in a DB and can serve out views of that data.
    * server starts up
    * subscribes to new blocks
    * https://lcd.technofractal.com/validatorsets/459777
    * Or `/stake/validators` every block with notifications and persist the validators present at that block
