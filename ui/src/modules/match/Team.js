import React, { useState } from 'react';
import PlayersList from '../player/PlayersList.js';
import './Team.css'

const Team = ({name, players=[], chanceOfWinning}) => {
    return (
        <div className="team-container">
            <div className="team-name">{name}</div> <br/>
            <div className="chance-of-winning">WIN Chance: <b>{chanceOfWinning}%</b></div>
            <PlayersList players={players} match={true}/>
        </div>
    );
}

export default Team;
