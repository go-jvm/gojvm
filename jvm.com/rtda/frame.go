package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
}




