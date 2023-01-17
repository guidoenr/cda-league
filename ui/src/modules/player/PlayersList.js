import React from 'react';
import PlayerCard from './PlayerCard'; // Import the PlayerCard component
import './PlayersList.css'

const PlayerList = ({players =[], small}) => {
    return (
        <div className="player-list">
                {/* Map over the players array and render a PlayerCard component for each player */}
            {players.map(player => (
                <PlayerCard player={player} small={small}/>
            ))}
        </div>
    )
};

export default PlayerList;
