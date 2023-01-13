import React, { useState } from 'react';
import PlayersList from '../player/PlayersList.js';

const Team = ({name, players=[], totalPlayers, chanceOfWinning}) => {
    return (
        <div className="team-container">
            <h4>{name}</h4>
            <h5>CHANCE OF WINNING:{chanceOfWinning}%</h5>
            <h6>Total players: {totalPlayers}</h6>
            <PlayersList players={players} />
        </div>
    );
}

export default Team;
