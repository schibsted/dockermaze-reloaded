# Story

### It’s year BD5. The Dockerbot is your only hope of survival in an unknown world. Will you be able to fix it in time?

Sirens sound. A synthetic voice says: "Emergency wake up, power is almost depleted, welcome Doctor to year BD5, your bodily functions should be repaired shortly. I regret to inform you that the cryogenic capsule has malfunctioned." 

You find yourself cold and alone at a much later time than you intended. You don't know what has become of the world out of your laboratory, nor would it be wise to try and find out by yourself. The emergency life support has kicked in and you have around three days until the power runs out to prepare yourself for whatever is expecting you in the open. 

The radiation hasn’t been fair to your Dockerbot, most parts have either suffered corruption or have automatically reverted to previous working versions due to it. Soon after executing its `build` and `run` sequences, you find that the `head` refuses to accept any of the parts which, in their current state, are unable to pass their verification tests. Its `legs` are unable to walk through paths, the `arms` no longer coordinate with the `head`, the `weapon` appears to be unable to distinguish allies, the `radio` has message integrity problems and, to wrap it all up, there seems to be some sort of connectivity problem between the `heart` and the `head`. 

With little time to waste, you place your identification token in the `DM2_TOKEN` enviroment variable and prepare yourself for a race for survival against the clock and the unknown dangers that may await you outside. How much of the robot will you be able to put together? Can its powers be a match for whatever futuristic foes inhabit this world?

**TL;DR: Your Dockerbot is broken, fix it. Download the challenge, get your identification token and survive!**

# Usage

All necessary files can be downloaded from this repository and run with Docker Compose.

```
git clone https://github.com/schibsted/dockermaze-reloaded
cd dockermaze-reloaded
docker-compose build
docker-compose up
```

In order for the Dockerbot to work, an identification token is necessary, which can be obtained from:

http://challenge.schibsted.com/

# Requirements

In order to run the challenge, the following Docker versions are needed:

- Docker Compose 1.6.0+
- Docker Engine 1.10.0+
