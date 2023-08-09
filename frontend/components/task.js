function Task({ task, onDelete, onUpdate }) {
    return (
        <div className="task">
            <h3>{task.title}</h3>
            <p>{task.description}</p>
            <input 
                type="checkbox" 
                checked={task.done}
                onChange={() => onUpdate({...task, done: !task.done})}
            /> Completed
            <button onClick={() => onDelete(task.id)}>Delete</button>
        </div>
    );
}

export default Task;
