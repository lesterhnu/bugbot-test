#include "cases/leaks.h"

#include <iostream>
#include <memory>

namespace leaks
{

    namespace
    {
        struct Node
        {
            std::shared_ptr<Node> next;
            int value;

            explicit Node(int v) : value(v) {}
            ~Node() { std::cout << "Node(" << value << ") destroyed\n"; }
        };
    } // namespace

    void shared_cycle_leak()
    {
        auto a = std::make_shared<Node>(1);
        auto b = std::make_shared<Node>(2);

        // Bug pattern: cyclic shared_ptr ownership prevents destruction.
        a->next = b;
        b->next = a;

        std::cout << "[shared_cycle_leak] created cycle between two nodes\n";
    }

} // namespace leaks
