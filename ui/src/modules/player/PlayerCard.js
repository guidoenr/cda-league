import React, { useState } from 'react';
import './PlayerCard.css'
import '../Util'
import {getStars, getPhoto, getColor} from "../Util";

const PlayerCard = ({ nickname, name, rank, elo, position}) => {

    // Use the useState hook to create a state variable to track the visibility of the player description
    const [isDescriptionVisible, setIsDescriptionVisible] = useState(false);


    return (
        <div className="player-card ">
            <img src={getPhoto(nickname)} alt={`${name}'s profile photo`} className="player-photo" />
            <div className="player-info">
                <h2 className="player-nickname">{nickname}</h2>
                <div className="player-rank">
                    {getStars(rank)}
                </div>
                <div className="player-name">{name}</div>
                <div className="player-elo">ELO: {elo}</div>
                <div className="position-shape" style={{backgroundColor: getColor(position)}}> {position.toUpperCase().substring(0,3)}</div>
                {/* Add a button to show and hide the player description */}
                {/* Only show the player description if the isDescriptionVisible state variable is true */}
            </div>
        </div>
    );
}

export default PlayerCard;