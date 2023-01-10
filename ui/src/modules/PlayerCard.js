import React, { useState } from 'react';
import './PlayerCard.css'

const PlayerCard = ({ photo, nickname, name, rank, elo, position}) => {
    const photoName = nickname.toLowerCase() + ".jpeg"
    photo = require(`./assets/cartoon/${photoName}`)

    // searching how many stars will the player have
    const renderStars = () => {
        let stars = [];
        for (let i = 0; i < rank; i++) {
            stars.push("★");
        }
        return stars;
    }

    // Use the useState hook to create a state variable to track the visibility of the player description
    const [isDescriptionVisible, setIsDescriptionVisible] = useState(false);

    // Define a function to toggle the visibility of the player description
    const toggleDescription = () => {
        setIsDescriptionVisible(!isDescriptionVisible);
    };

    // Set the color of the rectangle shape based on the value of the position prop
    // what a shit is this
    let shapeColor;
    switch (position) {
        case 'delantero':
            shapeColor = 'red';
            position = 'DEL'
            break;
        case 'defensor':
            shapeColor = 'blue';
            position = 'DEF'
            break;
        case 'volante':
            shapeColor = 'green'
            position = 'VOL'
            break
        default:
            shapeColor = '#333';
            position = 'UNK'
            break
    }

    return (
        <div className="player-card">
            <img src={photo} alt={`${name}'s profile photo`} className="player-photo" />
            <div className="player-info">
                <h2 className="player-nickname">{nickname}</h2>
                <div className="player-rank">
                    {renderStars()}
                </div>
                <div className="player-name">{name}</div>
                <div className="player-elo">ELO: {elo}</div>
                <div className="position-shape" style={{backgroundColor: shapeColor}}> {position.toUpperCase()}</div>
                {/* Add a button to show and hide the player description */}
                {/* Only show the player description if the isDescriptionVisible state variable is true */}
            </div>
        </div>
    );
}

export default PlayerCard;