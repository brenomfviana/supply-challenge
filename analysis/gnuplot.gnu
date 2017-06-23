set encoding iso_8859_1 
set grid 
set key top left 
set xlabel 'Quantidade de Produtores/Consumidores' 
set ylabel 'tempo (segundos)' 

# Etapas 50ms 
set title 'Resultado com Processamento durando 50ms' 
plot 'p1_50.txt' using ($2/1000):xticlabel(1) t 'Etapa 1' with linespoints
rep 'p2_50.txt' using ($2/1000):xticlabel(1) t 'Etapa 2' with linespoints
rep 'p3_50.txt' using ($2/1000):xticlabel(1) t 'Etapa 3' with linespoints
set terminal png font arial 20 size 1280,720
set output 'result_etapas_50ms.png' 
replot 

# Etapas 100ms 
set title 'Resultado com Processamento durando 50ms' 
plot 'p1_100.txt' using ($2/1000):xticlabel(1) t 'Etapa 1' with linespoints
rep 'p2_100.txt' using ($2/1000):xticlabel(1) t 'Etapa 2' with linespoints
rep 'p3_100.txt' using ($2/1000):xticlabel(1) t 'Etapa 3' with linespoints
set terminal png font arial 20 size 1280,720
set output 'result_etapas_100ms.png' 
replot 

# Etapa 1 FULL
set title 'Resultados da Etapa 1 variando o tempo de processamento do pedido' 
plot 'p1_50.txt' using ($2/1000):xticlabel(1) t '50ms' with linespoints
rep 'p1_100.txt' using ($2/1000):xticlabel(1) t '100ms' with linespoints
rep 'p1_500.txt' using ($2/1000):xticlabel(1) t '500ms' with linespoints
set terminal png font arial 20 size 1280,720
set output 'result_etapa1.png' 
replot 

# Etapa 2 FULL
set title 'Resultados da Etapa 2 variando o tempo de Processamento do pedido' 
plot 'p2_50.txt' using ($2/1000):xticlabel(1) t '50ms' with linespoints
rep 'p2_100.txt' using ($2/1000):xticlabel(1) t '100ms' with linespoints 
set terminal png font arial 20 size 1280,720
set output 'result_etapa2.png' 
replot 

# Etapa 3 FULL
set title 'Resultados da Etapa 3 variando o tempo de Processamento do pedido' 
plot 'p3_50.txt' using ($2/1000):xticlabel(1) t '50ms' with linespoints
rep 'p3_100.txt' using ($2/1000):xticlabel(1) t '100ms' with linespoints 
set terminal png font arial 20 size 1280,720
set output 'result_etapa3.png' 
replot 


