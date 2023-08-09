import { useState, useEffect } from 'react';
import TaskList from '../components/taskList';

function HomePage() {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        const fetchTasks = async () => {
            const response = await fetch("/api/tasks");
            const data = await response.json();
            setTasks(data);
        };
        fetchTasks();
    }, []);

    const handleDelete = async (id) => {
        await fetch(`/api/tasks`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ id })
        });
        setTasks(tasks.filter(task => task.id !== id));
    };

    const handleUpdate = async (updatedTask) => {
        await fetch(`/api/tasks`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedTask)
        });
        setTasks(tasks.map(task => task.id === updatedTask.id ? updatedTask : task));
    };

    return (
        <div className="container">
            <h1>Tasks</h1>
            <TaskList tasks={tasks} onDelete={handleDelete} onUpdate={handleUpdate} />
        </div>
    );
}

export default HomePage;
