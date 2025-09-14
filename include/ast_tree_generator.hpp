#include <iostream>
#include "antlr4-runtime.h"
#include "coralBaseListener.h"
#include "coralLexer.h"
#include "coralParser.h"

using namespace antlr4;
using namespace std;


namespace ASTNODES 
{
    struct ASTNode {
        virtual ~ASTNode() = default;  // virtual destructor for polymorphism
    };

    template <typename T>
    struct Expression {
        Expression* next;
        T value;
    };

    struct BlockNode : ASTNode {
        std::vector<std::shared_ptr<ASTNode>> statements;
    };

    struct LoopNode : ASTNode{
        int start_;
        int end_;
        int increment_;
        BlockNode* next;

    };

    template <typename T> 
    struct RewindNode : ASTNode {
        T key;
        BlockNode* next;
    };

    template <typename T>
    struct IfNode : ASTNode {
        T comparator1;
        T comparator2;
        string operator_;
        Expression<T> expression;
        BlockNode* next;
    };


    template <typename T> 
    struct BreakPointNode : ASTNode {
        T key;
        BlockNode* next;
    };

    template <typename T>
    struct AssignmentNode : ASTNode {
        string ID;
        Expression<T> expression;
        BlockNode* next;
    };

    

    // struct blockNode {
        
    // }



}