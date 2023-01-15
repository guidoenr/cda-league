import React, { useState } from 'react';
import PlayersList from '../player/PlayersList.js';
import './Team.css'

const Team = ({name, players=[], totalPlayers, chanceOfWinning}) => {
    return (
        <div className="team-container">
            <h4>{name}</h4>
            <h5>Chance of Winning: <b>{chanceOfWinning}%</b></h5>
            <PlayersList players={players} />
        </div>
    );
}

export default Team;
