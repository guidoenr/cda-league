import React, {useState} from 'react';
import './PlayerCard.css'
import '../Util'
import {getStars, getPhoto, getColor} from "../Util";
import Container from 'react-bootstrap/Container';

const PlayerCard = ({player, small} ) => {
    const [isSelected, setIsSelected] = useState(false);
    const handleClick = () => {
        setIsSelected(!isSelected);
    }

    const selected = isSelected ? "selected" : "";
    const className = small ? "player-card sm" : "player-card";

    return (
        <Container className={className + selected} onClick={handleClick}>
            <img src={getPhoto(player.nickname)} alt={`${player.name}'s profile photo`} className="photo" />
            <div className="info">
                <h2 className="nickname">{player.nickname}</h2>
                <div className="rank">
                    {getStars(player.rank)}
                </div>
                <div className="name">{player.name}</div>
                <div className="elo">ELO: {player.elo}</div>
                <div className="position-shape" style={{backgroundColor: getColor(player.position)}}> {player.position.toUpperCase().substring(0,3)}</div>
            </div>
        </Container>
    );
}

export default PlayerCard;