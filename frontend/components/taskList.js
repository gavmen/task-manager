import Task from './task';

function TaskList({ tasks, onDelete, onUpdate }) {
    return (
        <div className="taskList">
            {tasks.map(task => (
                <Task key={task.id} task={task} onDelete={onDelete} onUpdate={onUpdate} />
            ))}
        </div>
    );
}

export default TaskList;
