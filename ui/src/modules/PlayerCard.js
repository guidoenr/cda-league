import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faStar as fasStar } from '@fortawesome/free-solid-svg-icons';
import { faStar as farStar } from '@fortawesome/free-regular-svg-icons';
import './PlayerCard.css'

const PlayerCard = ({ photo, nickname, name, rank, position }) => {
    photo = require('./assets/cartoon/gonza.jpeg')

    const renderStars = () => {
        let stars = [];
        for (let i = 0; i < rank; i++) {
            stars.push("★");
        }
        return stars;
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
                <div className="player-rank-position">{position}</div>
            </div>
        </div>
    );
}

export default PlayerCard;