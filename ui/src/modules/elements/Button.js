import './Button.css'

const Button = ({textToDisplay}) => {
    return (
        <div className="button-container">
            <button className="button type1">
                {textToDisplay}
            </button>
        </div>
    )

}

export default Button;