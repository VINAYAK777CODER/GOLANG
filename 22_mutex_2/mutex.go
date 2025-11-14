package main
import ("fmt" 
"sync" )
type Post struct{
	veiws int
	mu sync.Mutex
}
func (p* Post) increment(wg *sync.WaitGroup){
	defer func ()  {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.veiws=p.veiws+1
	
}
func main(){
	var wg sync.WaitGroup
	// var mu sync.Mutex


	myPost:=Post{
		veiws:0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		 go myPost.increment(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.veiws,"is incremented value")
}