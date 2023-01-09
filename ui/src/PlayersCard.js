import React from 'react';
import './PlayersCard.css';

function PlayersCard(props) {
    const { players } = props;

    return (
        <div className="players-cards">
            {players.map(player => (
                <div key={player.ID} className="player-card">
                    <div
                        className="player-photo"
                        style={{ backgroundImage: 'url(./cartoon/cfacu.jpeg) no-repeat center/cover' }}

                    ></div>
                    <div className="player-info">
                        <h3>{player.nickname}</h3>

                        <p>{player.name}</p>
                        <p>{player.description}</p>
                        <p>
                            Age: {player.age} | Rank: {player.rank} | Position: {player.position}
                        </p>
                        <p>Elo: {player.elo}</p>
                        <p>
                            Goals per match: {player.goalsPerMatch} | Games won: {player.gamesWon}
                        </p>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default PlayersCard;
