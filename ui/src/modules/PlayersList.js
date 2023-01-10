import React, { useState, useEffect } from 'react';
import PlayerCard from './PlayerCard'; // Import the PlayerCard component

const PlayerList = () => {
    // Use the useState hook to create state variables for the player data and the loading state
    const [players, setPlayers] = useState([]);
    const [isLoading, setIsLoading] = useState(false);


    // Use the useEffect hook to fetch the player data from the REST API when the component mounts
    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/');
                const data = await response.json();
                setPlayers(data.players);
                setIsLoading(false);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers()
    }, []);

    return (
        <div>
            {isLoading ? (
                <div>Loading...</div>
            ) : (
                <div className="player-list">
                    {/* Map over the players array and render a PlayerCard component for each player */}
                    {players.map(player => (
                        <PlayerCard
                            key={player.id}
                            nickname={player.nickname}
                            name={player.name}
                            rank={player.rank}
                            position={player.position}
                            description={player.description}
                        />
                    ))}
                </div>
            )}
        </div>
    );
};

export default PlayerList;
