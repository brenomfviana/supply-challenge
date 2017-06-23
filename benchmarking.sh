#!/bin/bash

echo "Initializing Benchmarking..." &&
echo "---------------------"
cd "Phase_$1" &&
echo "Compiling Phase_$1..." &&
go build &&


echo "---------------------"
echo "Initial information..." &&
declare -A values &&
instances=(1 5 10 50 100 500 1000) &&
PRODUTORES=5000 &&

echo "Instances of Quantity of producers: ${instances[*]}" &&
echo "Instances of Quantity of consumers: ${instances[*]}" &&

echo "---------------------"
echo "Running..." &&

for ((i=0;i<7;i++)) do
	echo -n "Instance ${instances[$i]}:	"
	for ((j=0;j<10;j++)) do
		start=`date +%s%N`/1000000  &&
		./"Phase_$1" ${instances[$i]} ${instances[$i]}  >> saida.txt &&
		end=`date +%s%N`/1000000 &&
		runtime=$(((end-start))) &&
		values[$i,$j]=$runtime
		
		echo -n "$j "
	done
	echo ""
done

rm -f saida.txt
rm -f "Phase_$1" || rm -f "Phase_$1.exe"

echo "---------------------"
echo "Collected data in milliseconds"
# echo "${values[*]}"
for ((i=0;i<7;i++)) do
	echo -n "${instances[$i]}	"
    for ((j=0;j<10;j++)) do
    	echo -n "${values[$i,$j]}	"
    done
    echo ""
done

echo "---------------------"
declare -A medio
declare -A maior
declare -A menor
declare -A desvioP

echo "Inst	Average	Greater	Smaller	Variance"
for ((i=0;i<7;i++)) do
	menor[$i]=${values[$i,0]}
	maior[$i]=${values[$i,0]}
	
	for ((j=0;j<10;j++)) do
    	medio[$i]=$(( ${medio[$i]} + ${values[$i,$j]} ))

    	if [ "${values[$i,$j]}" -gt  "${maior[$i]}" ]; then
		   maior[$i]=${values[$i,$j]}
		fi
		if [ "${values[$i,$j]}" -lt "${menor[$i]}" ]; then
		   menor[$i]=${values[$i,$j]}
		fi

    done
    medio[$i]=$(( ${medio[$i]} / 10 ))

    variancia=0
    for ((j=0;j<10;j++)) do
    	v_m=$(( ${values[$i,$j]} - ${medio[$i]} ))
    	v_m_2=$(( $v_m * $v_m ))
    	variancia=$(( $variancia + $v_m_2 ))
    done
    variancia=$(( $variancia / ${instances[$i]} ))
    desvioP[$i]=$variancia
    
    echo "${instances[$i]}	${medio[$i]}	${maior[$i]}	${menor[$i]}	${desvioP[$i]}" 
done
