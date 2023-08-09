
let tasks = [
    { id: 1, title: "Sample Task 1", completed: false },
    { id: 2, title: "Sample Task 2", completed: true }
];

export default (req, res) => {
    const { method } = req;

    switch (method) {
        case 'GET':
            res.status(200).json(tasks);
            break;
        case 'POST':
            const newTask = {
                id: tasks.length + 1,
                title: req.body.title,
                completed: req.body.completed || false
            };
            tasks.push(newTask);
            res.status(201).json(newTask);
            break;
        case 'PUT':
            const taskId = req.body.id;
            const taskIndex = tasks.findIndex(task => task.id === taskId);
            if (taskIndex > -1) {
                tasks[taskIndex] = { ...tasks[taskIndex], ...req.body };
                res.status(200).json(tasks[taskIndex]);
            } else {
                res.status(404).json({ message: "Task not found" });
            }
            break;
        case 'DELETE':
            const idToDelete = parseInt(req.query.id);
            tasks = tasks.filter(task => task.id !== idToDelete);
            res.status(200).json({ message: "Task deleted" });
            break;
        default:
            res.setHeader('Allow', ['GET', 'POST', 'PUT', 'DELETE']);
            res.status(405).end(`Method ${method} Not Allowed`);
    }
};
