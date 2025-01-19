import { useNavigate, useParams } from "react-router-dom";
import TodoAppBar from "../components/TodoAppBar"

const Todo = () => {
  const { id } = useParams<{ id: string }>();

  const navigate = useNavigate();
  if (!id) {
    navigate('/login');
  }

  return (
    id ? <TodoAppBar id={id} /> : null
  );
}

export default Todo;