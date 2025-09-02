import { useState, useEffect } from 'react'
import './App.css'

const API_URL = 'http://localhost:1003'

function App() {
  const [todos, setTodos] = useState([])
  const [newTodo, setNewTodo] = useState('')
  const [filter, setFilter] = useState('all')
  const [loading, setLoading] = useState(false)
  const [operationLoading, setOperationLoading] = useState(false)

  const fetchTodos = async () => {
    try {
      setLoading(true)
      const response = await fetch(`${API_URL}/todo`)
      if (!response.ok) throw new Error('Помилка завантаження')
      const data = await response.json()
      setTodos(data || [])
    } catch (error) {
      console.error('Помилка завантаження:', error)
      alert('Не вдалося завантажити задачі')
      setTodos([])
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTodos()
  }, [])

  const addTodo = async () => {
    if (!newTodo.trim() || operationLoading) return

    try {
      setOperationLoading(true)
      const response = await fetch(`${API_URL}/todo`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          body: newTodo,
          completed: false
        })
      })

      if (!response.ok) throw new Error('Помилка додавання')
      
      const createdTodo = await response.json()
      setTodos(prev => [...(prev || []), createdTodo])
      setNewTodo('')
    } catch (error) {
      console.error('Помилка додавання:', error)
      alert('Не вдалося додати задачу')
    } finally {
      setOperationLoading(false)
    }
  }

  const toggleTodo = async (id, currentCompleted) => {
    if (operationLoading) return

    try {
      setOperationLoading(true)
      const response = await fetch(`${API_URL}/todo/${id}`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          completed: !currentCompleted
        })
      })

      if (!response.ok) throw new Error('Помилка оновлення')
      
      setTodos(prev => prev.map(todo =>
        todo.id === id ? { ...todo, completed: !currentCompleted } : todo
      ))
    } catch (error) {
      console.error('Помилка оновлення:', error)
      alert('Не вдалося оновити задачу')
    } finally {
      setOperationLoading(false)
    }
  }

  const deleteTodo = async (id) => {
    if (operationLoading) return
    if (!window.confirm('Видалити задачу?')) return

    try {
      setOperationLoading(true)
      const response = await fetch(`${API_URL}/todo/${id}`, {
        method: 'DELETE'
      })

      if (!response.ok) throw new Error('Помилка видалення')
      
      setTodos(prev => prev.filter(todo => todo.id !== id))
    } catch (error) {
      console.error('Помилка видалення:', error)
      alert('Не вдалося видалити задачу')
    } finally {
      setOperationLoading(false)
    }
  }

  const filteredTodos = (todos || []).filter(todo => {
    if (filter === 'completed') return todo.completed
    if (filter === 'active') return !todo.completed
    return true
  })

  return (
    <div className="app">
      <div className="container">
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
            disabled={operationLoading}
          />
          <button 
            onClick={addTodo} 
            className="add-btn" 
            disabled={operationLoading || !newTodo.trim()}
          >
            {operationLoading ? '⏳' : '➕'} Додати
          </button>
        </div>

        <div className="filters">
          <button 
            className={filter === 'all' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('all')}
            disabled={loading}
          >
            Всі
          </button>
          <button 
            className={filter === 'active' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('active')}
            disabled={loading}
          >
            Активні
          </button>
          <button 
            className={filter === 'completed' ? 'filter-btn active' : 'filter-btn'}
            onClick={() => setFilter('completed')}
            disabled={loading}
          >
            Виконані
          </button>
        </div>

        {loading ? (
          <div className="loading">Завантаження...</div>
        ) : (
          <div className="todo-list">
            {filteredTodos.length > 0 ? (
              filteredTodos.map(todo => (
                <div key={todo.id} className={`todo-item ${todo.completed ? 'completed' : ''}`}>
                  <label className="checkbox-label">
                    <input
                      type="checkbox"
                      checked={todo.completed}
                      onChange={() => toggleTodo(todo.id, todo.completed)}
                      className="checkbox"
                      disabled={operationLoading}
                    />
                    <span className="checkmark"></span>
                  </label>
                  <span className="todo-text">{todo.body}</span>
                  <button 
                    onClick={() => deleteTodo(todo.id)}
                    className="delete-btn"
                    title="Видалити"
                    disabled={operationLoading}
                  >
                    {operationLoading ? '⏳' : '🗑️'}
                  </button>
                </div>
              ))
            ) : (
              <div className="empty-state">
                <p>🎉 Немає задач! Додайте першу задачу!</p>
              </div>
            )}
          </div>
        )}

        <div className="stats">
          <span>Залишилось задач: {(todos || []).filter(t => !t.completed).length}</span>
          <span>Всього задач: {(todos || []).length}</span>
        </div>
      </div>
    </div>
  )
}

export default App