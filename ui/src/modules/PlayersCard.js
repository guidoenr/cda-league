import React, {useEffect, useState} from 'react';
import './PlayersCard.css';

function PlayersCard(props) {
    const { players } = props;
    return (
        <div className="players-cards">
            {players.map(player => {
                const photo = require(`./assets/cartoon/${player.nickname.toLowerCase()}.jpeg`); // import the photo using the player's name

                return (
                    <div key={player.ID} className="player-card" data-position={player.position}>
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

function PlayersCardsList() {
    const [players, setPlayers] = useState([]);

    useEffect(() => {
        // Fetch players data from API
        fetch('http://localhost:8080/players/rank')
            .then(response => response.json())
            .then(data => setPlayers(data.players));
    }, []);

    return (
        <div className="center">
            <h2>Players Cards</h2>
            <PlayersCard players={players} />
        </div>
    );
}

export default PlayersCardsList;
