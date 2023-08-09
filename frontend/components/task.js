function Task({ task, onDelete, onUpdate }) {
    return (
        <div className="task">
            <h3>{task.title}</h3>
            <p>{task.description}</p>
            <div>
            <input 
                type="checkbox" 
                checked={task.done}
                onChange={() => onUpdate({...task, done: !task.done})}
            /> <span>Completed</span>
            </div>
            <button onClick={() => onDelete(task.id)}>Delete</button>
        </div>
    );
}

export default Task;
