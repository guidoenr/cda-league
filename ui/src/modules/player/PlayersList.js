import React from 'react';
import PlayerCard from './PlayerCard'; // Import the PlayerCard component
import './PlayersList.css'

const PlayerList = ({players}) => {
    console.log("PLAYERLIST" + players)
    return (
        <div className="player-list">
                {/* Map over the players array and render a PlayerCard component for each player */}
            {players.map(player => (
                <PlayerCard
                    key={player.ID}
                    nickname={player.nickname}
                    name={player.name}
                    rank={player.rank}
                    position={player.position}
                    description={player.description}
                    elo={player.elo}
                />
            ))}
        </div>
    )
};

export default PlayerList;
