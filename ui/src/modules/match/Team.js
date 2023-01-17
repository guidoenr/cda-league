import React  from 'react';
import PlayersList from '../player/PlayersList.js';
import {getTeamLogo} from "../Util";
import './Team.css'

const Team = ({name, players=[], chanceOfWinning}) => {
    const logo = getTeamLogo(name)

    return (
        <div className="team-container">
            <div><img src={logo} className="team-logo"  alt={name}/></div>
            <div className="team-name">{name}</div> <br/>
            <div className="chance-of-winning">WIN Chance: <b>{chanceOfWinning}%</b></div>
            <PlayersList players={players} match={true}/>
        </div>
    );
}

export default Team;
