#!/bin/bash
set -e 

array_name=(Class Fieldref Type InterfaceMethodref String Integer Float Double NameAndType Utf8 MethodHandle MethodType InvokeDynamic) 
    # CONSTANT_Class = 7

	# CONSTANT_Fieldref = 9

	# CONSTANT_Type = 10

	# CONSTANT_InterfaceMethodref = 11

	# CONSTANT_String = 8

	# CONSTANT_Integer = 3

	# CONSTANT_Float = 4

	# CONSTANT_Long = 5

	# CONSTANT_Double = 6

	# CONSTANT_NameAndType = 12

	# CONSTANT_Utf8 = 1

	# CONSTANT_MethodHandle = 15

	# CONSTANT_MethodType = 16

	# CONSTANT_InvokeDynamic = 18

# 1. 生成15句case语句

# if [ -e "tmp" ];
# then
#      rm -f ./tmp 
# fi

# for c in ${array_name[@]}
# do
#     printf 'case CONSTANT_%s:return &Constant%sInfo{}\n' $c $c  >> tmp
#     let "ii++"
# done 


# mapping，太麻烦，砍掉
# t_Integer="int32" 
# t_Float="float32"
# t_Long="int64"
# t_Double="float64" 


FILE=cp_numeric.go
PACKAGE=classfile
cp_numerics=(Integer Float Long Double)
# for c in ${cp_numerics[@]}
# do 

#     printf '%s %s\n' $c $t_$c
# done 


# 生成 cp_numeric.go
# if [ -e "cp_numeric.go" ]
# then 
#     rm -f ./$FILE
# fi
# # 2.生成 cp_numeric
# printf 'package %s\n' $PACKAGE >> $FILE
# for c in ${cp_numerics[@]}
# do
#     # printf '%s %s\n' $c 
    
#     printf '
# type Constant%sInfo struct {
# 	val XXX
# }

# func (self *Constant%sInfo) readInfo(reader *ClassReader) {
# 	bytes := reader.readUint32()
# 	//self.val = math.Float32frombits()
# }\n' $c $c >> $FILE
# done 

# 生成 constant_numeric.go





