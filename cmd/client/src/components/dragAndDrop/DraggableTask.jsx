import {useDraggable} from '@dnd-kit/react';

export default function DraggableTask({ task, children }) {
    const { ref } = useDraggable({
        id: task.id
    });

    return (
        <div ref={ref}>
            {children}
        </div>
    )
}