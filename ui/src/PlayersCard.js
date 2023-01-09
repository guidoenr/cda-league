import React, { useState } from 'react';
import './PlayersCard.css';

function PlayersCard(props) {
    const { players } = props;

    return (
        <div className="players-cards">
            {players.map(player => {
                const photo = require(`./cartoon/${player.nickname.toLowerCase()}.jpeg`); // import the photo using the player's name

                return (
                    <div key={player.ID} className="player-card">
                        <div
                            className="player-photo"
                            style={{ backgroundImage: `url(${photo})` }}
                        ></div>
                        <div className="player-info">
                            <h3>{player.nickname}</h3>
                            <p>
                <span className="rank-stars">
                  {Array.from({ length: player.rank }, () => '⭐️').join('')}
                </span>
                            </p>
                            <p>{player.name}</p>
                            <p>{player.position}</p>
                            <p>
                                <b>ELO</b>:{player.elo}
                            </p>
                        </div>
                    </div>
                );
            })}
        </div>
    );
}

export default PlayersCard;
