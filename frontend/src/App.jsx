import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [todos, setTodos] = useState([])
  const [newTodo, setNewTodo] = useState('')
  const [filter, setFilter] = useState('all')

  useEffect(() => {
    const mockTodos = [
      { id: 1, body: 'Вивчити React', completed: false },
      { id: 2, body: 'Написати Todo App', completed: true },
      { id: 3, body: 'Підключити бекенд', completed: false }
    ]
    setTodos(mockTodos)
  }, [])

  const filteredTodos = todos.filter(todo => {
    if (filter === 'completed') return todo.completed
    if (filter === 'active') return !todo.completed
    return true
  })

  const addTodo = () => {
    if (newTodo.trim()) {
      setTodos([...todos, {
        id: Date.now(),
        body: newTodo,
        completed: false
      }])
      setNewTodo('')
    }
  }

  const toggleTodo = (id) => {
    setTodos(todos.map(todo =>
      todo.id === id ? { ...todo, completed: !todo.completed } : todo
    ))
  }

  const deleteTodo = (id) => {
    setTodos(todos.filter(todo => todo.id !== id))
  }

  return (
    <div className="app">
      <div className="container">
        {/* Заголовок */}
        <header className="header">
          <h1>✅ Мої Завдання</h1>
          <p>Керуй своїми справами</p>
        </header>

        <div className="input-section">
          <input
            type="text"
            value={newTodo}
            onChange={(e) => setNewTodo(e.target.value)}
            placeholder="Що потрібно зробити?..."
            onKeyPress={(e) => e.key === 'Enter' && addTodo()}
            className="todo-input"
          />
          <button onClick={addTodo} className="add-btn">
            ➕ Додати
          </button>
        </div>

        <div className="filters">
          <button 
            className={filter === 'all' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('all')}
          >
            Всі
          </button>
          <button 
            className={filter === 'active' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('active')}
          >
            Активні
          </button>
          <button 
            className={filter === 'completed' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('completed')}
          >
            Виконані
          </button>
        </div>

        <div className="todo-list">
          {filteredTodos.map(todo => (
            <div key={todo.id} className={`todo-item ${todo.completed ? 'completed' : ''}`}>
              <label className="checkbox-label">
                <input
                  type="checkbox"
                  checked={todo.completed}
                  onChange={() => toggleTodo(todo.id)}
                  className="checkbox"
                />
                <span className="checkmark"></span>
              </label>
              <span className="todo-text">{todo.body}</span>
              <button 
                onClick={() => deleteTodo(todo.id)}
                className="delete-btn"
                title="Видалити"
              >
                🗑️
              </button>
            </div>
          ))}
        </div>

        <div className="stats">
          <span>Залишилось задач: {todos.filter(t => !t.completed).length}</span>
        </div>
      </div>
    </div>
  )
}

export default App