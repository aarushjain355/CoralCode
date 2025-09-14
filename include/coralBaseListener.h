
// Generated from coral.g4 by ANTLR 4.13.2

#pragma once


#include "antlr4-runtime.h"
#include "coralListener.h"
#include "ast_tree_generator.hpp"

using namespace std;
using namespace ASTNODES;

/**
 * This class provides an empty implementation of coralListener,
 * which can be extended to create a listener which only needs to handle a subset
 * of the available methods.
 */
class  coralBaseListener : public coralListener {
public:

  std::shared_ptr<BlockNode> root = std::make_shared<BlockNode>();
  std::stack<std::shared_ptr<BlockNode>> blockStack;
  bool root_node_entered_ = false;
  virtual void enterCommand_first(coralParser::Command_firstContext * /*ctx*/) override { }
  virtual void exitCommand_first(coralParser::Command_firstContext * /*ctx*/) override { }

  virtual void enterArg_first(coralParser::Arg_firstContext * /*ctx*/) override { }
  virtual void exitArg_first(coralParser::Arg_firstContext * /*ctx*/) override { }

  virtual void enterCommand_second(coralParser::Command_secondContext * /*ctx*/) override { }
  virtual void exitCommand_second(coralParser::Command_secondContext * /*ctx*/) override { }

  virtual void enterArg_second(coralParser::Arg_secondContext * /*ctx*/) override { }
  virtual void exitArg_second(coralParser::Arg_secondContext * /*ctx*/) override { }

  virtual void enterLoop(coralParser::LoopContext * /*ctx*/) override { }
  virtual void exitLoop(coralParser::LoopContext *ctx) override { 
    auto loop_node_ = std::make_shared<LoopNode>();
    loop_node_->start_ = std::stoi(ctx->INTEGER(0)->getText());
    loop_node_->end_ = std::stoi(ctx->INTEGER(1)->getText());
    loop_node_->increment_ = std::stoi(ctx->INTEGER(2)->getText());
    root->statements.push_back(loop_node_);
    if (root_node_entered_ == false) 
    {
      
      root_node_entered_ = true;
      //blockStack.emplace(loop_node);
    }
    

    //root_copy->statements.push_back(loop_node_);
  }

  virtual void enterBlock(coralParser::BlockContext * /*ctx*/) override { }
  virtual void exitBlock(coralParser::BlockContext * /*ctx*/) override { }

  virtual void enterStatement(coralParser::StatementContext * /*ctx*/) override { }
  virtual void exitStatement(coralParser::StatementContext * /*ctx*/) override { }

  virtual void enterIfStatement(coralParser::IfStatementContext * /*ctx*/) override { }
  virtual void exitIfStatement(coralParser::IfStatementContext * /*ctx*/) override { }

  virtual void enterCondition(coralParser::ConditionContext * /*ctx*/) override { }
  virtual void exitCondition(coralParser::ConditionContext * /*ctx*/) override { }

  virtual void enterAssignment(coralParser::AssignmentContext * /*ctx*/) override { }
  virtual void exitAssignment(coralParser::AssignmentContext * /*ctx*/) override { }

  virtual void enterFloat(coralParser::FloatContext * /*ctx*/) override { }
  virtual void exitFloat(coralParser::FloatContext * /*ctx*/) override { }

  virtual void enterMulDiv(coralParser::MulDivContext * /*ctx*/) override { }
  virtual void exitMulDiv(coralParser::MulDivContext * /*ctx*/) override { }

  virtual void enterAddSub(coralParser::AddSubContext * /*ctx*/) override { }
  virtual void exitAddSub(coralParser::AddSubContext * /*ctx*/) override { }

  virtual void enterParens(coralParser::ParensContext * /*ctx*/) override { }
  virtual void exitParens(coralParser::ParensContext * /*ctx*/) override { }

  virtual void enterANOTHERVARIABLE(coralParser::ANOTHERVARIABLEContext * /*ctx*/) override { }
  virtual void exitANOTHERVARIABLE(coralParser::ANOTHERVARIABLEContext * /*ctx*/) override { }

  virtual void enterString(coralParser::StringContext * /*ctx*/) override { }
  virtual void exitString(coralParser::StringContext * /*ctx*/) override { }

  virtual void enterInt(coralParser::IntContext * /*ctx*/) override { }
  virtual void exitInt(coralParser::IntContext * /*ctx*/) override { }

  virtual void enterProgram(coralParser::ProgramContext * /*ctx*/) override { }
  virtual void exitProgram(coralParser::ProgramContext * /*ctx*/) override { }


  virtual void enterEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void exitEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void visitTerminal(antlr4::tree::TerminalNode * /*node*/) override { }
  virtual void visitErrorNode(antlr4::tree::ErrorNode * /*node*/) override { }

};

