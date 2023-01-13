import React, { useState } from 'react';
import PlayersList from '../player/PlayersList.js';

const Team = ({name, players, totalPlayers, chanceOfWinning}) => {
    return (
        <div className="team-container">
            <h2>{name}</h2>
            <h3>CHANCE OF WINNING:{chanceOfWinning}%</h3>
            <h4>Total players: {totalPlayers}</h4>
            <PlayersList players={players} />
        </div>
    );
}

export default Team;
