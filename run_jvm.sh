###
 # @Author: XPectuer
 # @LastEditor: XPectuer
###
JVM_GO=./bin/jvmgo
__CLASS_PATH='/Users/alex/projects/go_projs/jvm/test/java/'  
__JRE_PATH='/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre' 
__TEST_CLASS=com.zjut.test.Test
$JVM_GO -Xjre $__JRE_PATH -cp $__CLASS_PATH $__TEST_CLASS  