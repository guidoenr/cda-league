import React from 'react';
import './PlayerCard.css'
import '../Util'
import {getStars, getPhoto, getColor} from "../Util";

const PlayerCard = ({player}) => {

    return (
        <div className="player-card ">
            <img src={getPhoto(player.nickname)} alt={`${player.name}'s profile photo`} className="player-photo" />
            <div className="player-info">
                <h2 className="player-nickname">{player.nickname}</h2>
                <div className="player-rank">
                    {getStars(player.rank)}
                </div>
                <div className="player-name">{player.name}</div>
                <div className="player-elo">ELO: {player.elo}</div>
                <div className="position-shape" style={{backgroundColor: getColor(player.position)}}> {player.position.toUpperCase().substring(0,3)}</div>
            </div>
        </div>
    );
}

export default PlayerCard;